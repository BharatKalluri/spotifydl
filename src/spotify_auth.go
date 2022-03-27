package spotifydl

import (
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
)

// UserData is a struct to hold all variables
type UserData struct {
	UserClient      *spotify.Client
	TrackList       []spotify.FullTrack
	SimpleTrackList []spotify.SimpleTrack
	YoutubeIDList   []string
}

// InitAuth starts Authentication
func InitAuth() *spotify.Client {
	ctx := context.Background()
	// Please do not misuse :)
	config := &clientcredentials.Config{
		ClientID:     "07d728d8751646219ab0532ca21ba982",
		ClientSecret: "6ad82c4fd7cc498fbf738ea08f4135d3",
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)

	return client
}
