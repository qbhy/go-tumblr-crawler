package response

import "fmt"

type Photo struct {
	Tumblr
	Posts PhotoPosts `xml:"posts"`
}

type PhotoPosts struct {
	BasePosts
	Post []PhotoPost `xml:"post"`
}

type PhotoPost struct {
	BasePost
	Photos []string `xml:"photo-url"`
}

func (post *PhotoPost) ParsePhotosUrl() map[string]string {
	photos := make(map[string]string)

	fmt.Println("图片数量:" + string(len(post.Photos)))

	for key, photoUrl := range post.Photos {
		photos[post.Tumblelog.Title+string(key)] = photoUrl
		fmt.Println(photoUrl)
	}

	return photos
}

// photo-url
