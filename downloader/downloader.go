package downloader

import (
	"tumblr-crawler/downloader/response"
	"fmt"
	"github.com/qbhy/go-utils"
	"github.com/parnurzeal/gorequest"
	"path"
)

var Channel chan bool
var request *gorequest.SuperAgent

func init() {
	Channel = make(chan bool)
	request = gorequest.New()
}

func downloadVideos(site *Site, posts *response.VideoPosts) {
	for index, post := range posts.Post {
		fmt.Println("index:", index)
		go downloadFile(site.ProxyConfig, post.ParseVideoUrl(), site.videoPath, post.Tumblelog.Title+string(index)+".mp4")
	}
}

func downloadPhotos(site *Site, posts *response.PhotoPosts) {
	for _, post := range posts.Post {
		for index, url := range post.ParsePhotosUrl() {
			fmt.Println("index:", index)
			go downloadFile(site.ProxyConfig, url, site.photoPath, post.Tumblelog.Title+string(index))
		}
	}
}

func downloadFile(proxy ProxyConfig, url string, filePath string, filename string) {

	realPath := path.Join(filePath, filename)

	if exists, _ := utils.PathExists(realPath); exists {
		return
	}

	res, body, err := request.Proxy(proxy.Https).Get(url).End()

	if err != nil {
		fmt.Println(err, body, res)
		return
	}

	utils.FilePutContents(realPath, []byte(body))
}
