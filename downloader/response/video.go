package response

type Video struct {
	Tumblr
	Posts VideoPosts `xml:"posts"`
}

type VideoPosts struct {
	BasePosts
	Post []VideoPost `xml:"post"`
}

type VideoPost struct {
	BasePost
	VideoPlayer string `xml:"video-player"`
}

func (post *VideoPost) ParseVideoUrl() string {
	return post.VideoPlayer
}
