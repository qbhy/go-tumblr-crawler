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

type VideoPosts struct {
	BasePosts
	Post []VideoPost `xml:"post"`
}

type VideoPost struct {
	BasePost
	VideoPlayer string `xml:"video-player"`
}

func (post *VideoPost) ParseVideoUrl() string {

	player := VideoPlayer{}

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

		panic(err)
	}

	return player.Source.Src
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
