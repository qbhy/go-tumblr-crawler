package response

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Video struct {
	Tumblr
	Posts VideoPosts `xml:"posts"`
}

func NewVideo() *Video {
	return &Video{}
}

type VideoPosts struct {
	BasePosts
	Post []VideoPost `xml:"post"`
}

type VideoPost struct {
	BasePost
	VideoPlayer string `xml:"video-player"`
}

func (post *VideoPost) ParseVideoUrl() (bool, string) {

	player := VideoPlayer{}

	fmt.Println("post.VideoPlayer", post.VideoPlayer)

	if strings.Contains(post.VideoPlayer, "instagram-media") {
		return false, ""
	}

	playerString := strings.Replace(post.VideoPlayer, "&lt;", "<", -1)
	playerString = strings.Replace(playerString, "&gt;", ">", -1)
	playerString = strings.Replace(playerString, "'", "\"", -1)
	playerString = strings.Replace(playerString, "\r<", "<", -1)
	playerString = strings.Replace(playerString, "\n<", "<", -1)
	playerString = strings.Replace(playerString, ">\r", ">", -1)
	playerString = strings.Replace(playerString, ">\n", ">", -1)
	playerString = strings.Replace(playerString, "{\"", "{'", -1)
	playerString = strings.Replace(playerString, "\"}", "'}", -1)
	playerString = strings.Replace(playerString, "\":", "':", -1)
	playerString = strings.Replace(playerString, ",\"", ",'", -1)
	playerString = strings.Replace(playerString, "\",", "',", -1)
	playerString = strings.Replace(playerString, ":\"", ":'", -1)
	playerString = strings.Replace(playerString, "\">", "\"/>", -1)
	playerString = strings.Replace(playerString, "muted data-crt-video ", " ", -1)

	err := xml.Unmarshal([]byte(playerString), &player)

	if err != nil {
		fmt.Println(err)

		fmt.Println("playerString", playerString)

		panic(err)
	}

	return true, player.Source.Src
}

type VideoPlayer struct {
	XMLName xml.Name    `xml:"video"`
	Data    string      `xml:"data-crt-options"`
	Source  VideoSource `xml:"source"`
}

type VideoSource struct {
	Src  string `xml:"src,attr"`
	Type string `xml:"type,attr"`
}
