package downloader

import (
	"tumblr-crawler/downloader/response"
	"fmt"
	"github.com/qbhy/go-utils"
	"github.com/parnurzeal/gorequest"
	"path"
	"sync"
	"strings"
	"path/filepath"
	"tumblr-crawler/config"
)

var request *gorequest.SuperAgent
var WaitGroupInstance sync.WaitGroup

func init() {
	request = gorequest.New()
}

func downloadVideos(site *Site, posts response.VideoPosts) {
	for _, post := range posts.Post {
		if ok, videoUrl := post.ParseVideoUrl(); ok {
			filename := filepath.Base(videoUrl) + ".mp4"
			filename = strings.Replace(filename, " ", "", -1)

			WaitGroupInstance.Add(1)
			go downloadFile(site.ProxyConfig, videoUrl, site.videoPath, filename)
		}
	}
}

func downloadPhotos(site *Site, posts response.PhotoPosts) {
	for _, post := range posts.Post {
		for _, url := range post.ParsePhotosUrl() {
			WaitGroupInstance.Add(1)
			go downloadFile(site.ProxyConfig, url, site.photoPath, filepath.Base(url))
		}
	}
}

func downloadFile(proxy config.ProxyConfig, url string, filePath string, filename string) {

	realPath := path.Join(filePath, filename)

	if exists, _ := utils.PathExists(realPath); exists {
		defer WaitGroupInstance.Done()
		return
	}

	res, body, err := request.Proxy(proxy.Https).Get(url).End()

	if err != nil {
		fmt.Println("下载失败:", url, err, body, res)
		defer WaitGroupInstance.Done()
		return
	}

	utils.FilePutContents(realPath, []byte(body))
	defer WaitGroupInstance.Done()
}
