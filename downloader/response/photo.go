package response

type Photo struct {
	Tumblr
	Posts PhotoPosts `xml:"posts"`
}

func NewPhoto() *Photo {
	return &Photo{}
}

type PhotoPosts struct {
	BasePosts
	Post []PhotoPost `xml:"post"`
}

type PhotoPost struct {
	BasePost
	Photos   []string `xml:"photo-url"`
	PhotoSet PhotoSet `xml:"photoset"`
}

type PhotoSet struct {
	Photo []PhotoSetItem `xml:"photo"`
}

type PhotoSetItem struct {
	Photos []string `xml:"photo-url"`
}

func (post *PhotoPost) ParsePhotosUrl() map[string]string {
	photos := make(map[string]string)

	if len(post.PhotoSet.Photo) > 0 {
		for key, photo := range post.PhotoSet.Photo {
			photos[post.Tumblelog.Title+string(key)] = photo.Photos[0]
		}
	} else {
		photos[post.Tumblelog.Title+"cover"] = post.Photos[0]
	}

	return photos
}

// photo-url
