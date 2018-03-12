package spotifydl

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const developerKey = "AIzaSyCkeEYBmBP5pfyjhYf8QHCRnfGQRsBM5hE"

func GetYoutubeIds(songName string) string {
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	call := service.Search.List("id,snippet").
		Q(songName)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	videos := make(map[string]string)
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}
	var videosIds []string
	for id := range videos {
		videosIds = append(videosIds, id)
	}
	return videosIds[0]
}
