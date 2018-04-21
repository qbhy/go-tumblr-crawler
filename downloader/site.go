package downloader

import (
	"github.com/qbhy/go-utils"
	"path"
	"os"
	"strings"
	"github.com/parnurzeal/gorequest"
	"fmt"
	"encoding/xml"
	response2 "tumblr-crawler/downloader/response"
)

func NewSite(site string, config ProxyConfig) *Site {
	return &Site{
		Site:        site,
		ProxyConfig: config,
	}
}

var BaseUrl = "http://{site}.tumblr.com/api/read?type={mediaType}&num={num}&start={start}"

func GenerateMediaUrl(site string, mediaType string, num string, start string) string {
	mediaUrl := strings.Replace(BaseUrl, "{site}", site, -1)
	mediaUrl = strings.Replace(mediaUrl, "{mediaType}", mediaType, -1)
	mediaUrl = strings.Replace(mediaUrl, "{num}", num, -1)
	mediaUrl = strings.Replace(mediaUrl, "{start}", start, -1)
	return mediaUrl
}

type Site struct {
	Site        string
	ProxyConfig ProxyConfig
	currentPath string
	sitePath    string
	videoPath   string
	photoPath   string
	request     *gorequest.SuperAgent
}

func (this *Site) StartDownload() {
	this.init()
	this.DownloadVideo()
	//this.DownloadPhoto()
}

func (this *Site) init() {
	this.currentPath = path.Join(utils.CurrentPath(), "files")

	if exists, _ := utils.PathExists(this.currentPath); !exists {
		os.Mkdir(this.currentPath, 0755)
	}

	this.sitePath = path.Join(this.currentPath, this.Site)
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

func (this *Site) DownloadMedia(mediaType string, start int) {

	mediaUrl := GenerateMediaUrl(this.Site, mediaType, "50", string(start))
	res, responseString, err := this.request.Get(mediaUrl).End()
	fmt.Println("mediaUrl", mediaUrl)

	if err != nil && res.StatusCode == 404 {
		fmt.Println(res)
		fmt.Println(err)
		fmt.Println("site does not exist", this.Site)
		return
	}

	if mediaType == "video" {
		video := response2.Video{}
		err := xml.Unmarshal([]byte(responseString), &video)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		go downloadVideos(this, &video.Posts)
	} else {
		photo := response2.Photo{}
		err := xml.Unmarshal([]byte(responseString), &photo)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		go downloadPhotos(this, &photo.Posts)
	}
}
