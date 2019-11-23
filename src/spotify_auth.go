package spotifydl

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"

	// transport is used as dependency for youtube API
	_ "google.golang.org/api/googleapi/transport"
	_ "google.golang.org/api/youtube/v3"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

// UserData is a struct to hold all variables
type UserData struct {
	UserClient    spotify.Client
	TrackList     []spotify.FullTrack
	YoutubeIDList []string
}

// InitAuth starts Authentication
func InitAuth() spotify.Client {
	// Please do not misuse :)
	config := &clientcredentials.Config{
		ClientID:     "07d728d8751646219ab0532ca21ba982",
		ClientSecret: "6ad82c4fd7cc498fbf738ea08f4135d3",
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	return client
}
