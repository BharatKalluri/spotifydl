package spotifydl

import (
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const developerKey = "AIzaSyCqHX9XaJgLbUF_qdSVVY67fKCD6Okqa0U"

// GetYoutubeIds takes the query as string and returns the search results video ID's
func GetYoutubeIds(songName string) (string, bool) {
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
	if len(videosIds) == 0 {
		return "", false
	}
	return videosIds[0], true
}
