package spotifydl

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// Please do not misuse :)
const developerKey = "AIzaSyDQn4VAc4MzrKOjo2sv5ucmKsQUIfKFaSE"

// GetYoutubeIds takes the query as string and returns the search results video ID's
func GetYoutubeIds(songName string) string {
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	// Video category ID 10 is for music videos
	call := service.Search.List("id,snippet").Q(songName).VideoCategoryId("10").Type("video")
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			return item.Id.VideoId
		}
	}
	// TODO: Handle when the query returns no songs (highly unlikely since the query is coming from spotify though)
	return ""
}
