package response

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

func (post *PhotoPost) ParsePhotosUrl() *map[string]string {
	photos := make(map[string]string)

	for key, photoUrl := range post.Photos {
		photos[post.Tumblelog.Title+string(key)] = photoUrl
	}

	return &photos
}

// photo-url
