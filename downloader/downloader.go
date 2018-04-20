package downloader

import (
	"tumblr-crawler/downloader/response"
	"fmt"
)

func downloadVideos(posts *response.VideoPosts) {
	for _, post := range posts.Post {
		fmt.Println(post.ParseVideoUrl())
	}
}

func downloadPhotos(posts *response.PhotoPosts) {
	for _, post := range posts.Post {
		fmt.Println(post.ParsePhotosUrl())
	}
}
