package downloader

import (
	"encoding/xml"
	"fmt"
	"math/big"
	"os"
	"path"
	"strings"

	"github.com/qbhy/go-tumblr-crawler/config"

	"github.com/parnurzeal/gorequest"
	response2 "github.com/qbhy/go-tumblr-crawler/downloader/response"
	"github.com/qbhy/go-utils"
)

func NewSite(site config.SiteConfig, config config.ProxyConfig) *Site {
	return &Site{
		Site:        site,
		ProxyConfig: config,
	}
}

const (
	BaseUrl    = "http://{site}.tumblr.com/api/read?type={mediaType}&num={num}&start={start}"
	PageNumber = 50
)

func GenerateMediaUrl(site string, mediaType string, num int64, start int64) string {
	mediaUrl := strings.Replace(BaseUrl, "{site}", site, -1)
	mediaUrl = strings.Replace(mediaUrl, "{mediaType}", mediaType, -1)
	mediaUrl = strings.Replace(mediaUrl, "{num}", big.NewInt(num).String(), -1)
	mediaUrl = strings.Replace(mediaUrl, "{start}", big.NewInt(start).String(), -1)
	fmt.Println("site:", site)
	fmt.Println("mediaType:", mediaType)
	fmt.Println("num:", num)
	fmt.Println("start:", start)
	fmt.Println("mediaUrl:", mediaUrl)
	return mediaUrl
}

type Site struct {
	Site        config.SiteConfig
	ProxyConfig config.ProxyConfig
	currentPath string
	sitePath    string
	videoPath   string
	photoPath   string
	request     *gorequest.SuperAgent
}

func (this *Site) StartDownload() {
	this.Init()

	if this.Site.Video {
		WaitGroupInstance.Add(1)
		go this.DownloadVideo()
	}
	if this.Site.Photo {
		WaitGroupInstance.Add(1)
		go this.DownloadPhoto()
	}
}

func (this *Site) Init() {
	this.currentPath = path.Join(utils.CurrentPath(), "files")

	if exists, _ := utils.PathExists(this.currentPath); !exists {
		os.Mkdir(this.currentPath, 0755)
	}

	this.sitePath = path.Join(this.currentPath, this.Site.Site)
	this.request = gorequest.New().Proxy(this.ProxyConfig.Https)

	if exists, _ := utils.PathExists(this.sitePath); !exists {
		os.Mkdir(this.sitePath, 0755)
	}
}

func (this *Site) DownloadVideo() {
	this.videoPath = path.Join(this.sitePath, "video")
	if exists, _ := utils.PathExists(this.videoPath); !exists {
		os.Mkdir(this.videoPath, 0755)
	}
	this.DownloadMedia("video", 0)
}

func (this *Site) DownloadPhoto() {
	this.photoPath = path.Join(this.sitePath, "photo")
	if exists, _ := utils.PathExists(this.photoPath); !exists {
		os.Mkdir(this.photoPath, 0755)
	}
	this.DownloadMedia("photo", 0)
}

func (this *Site) DownloadMedia(mediaType string, start int64) {

	for {

		mediaUrl := GenerateMediaUrl(this.Site.Site, mediaType, PageNumber, start)

		res, responseString, err := this.request.Get(mediaUrl).End()
		fmt.Println("start: ", start)
		fmt.Println("mediaUrl: ", mediaUrl)

		if err != nil || res.StatusCode == 404 {
			fmt.Println(res)
			fmt.Println(err)
			fmt.Println("site does not exist", this.Site)
			break
		}

		if mediaType == "video" {
			video := response2.NewVideo()
			err := xml.Unmarshal([]byte(responseString), &video)
			if err != nil {
				fmt.Printf("error: %v", err)
				break
			} else if len(video.Posts.Post) <= 0 {
				fmt.Println("没有更多内容了")
				break
			}

			downloadVideos(this, video.Posts)
		} else {
			photo := response2.NewPhoto()
			err := xml.Unmarshal([]byte(responseString), &photo)
			if err != nil {
				fmt.Printf("error: %v", err)
				break
			} else if len(photo.Posts.Post) <= 0 {
				fmt.Println("没有更多内容了")
				break
			}

			downloadPhotos(this, photo.Posts)
		}
		start += PageNumber
	}

	defer WaitGroupInstance.Done()
}
