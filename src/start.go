package spotifydl

import (
	"fmt"

	"github.com/zmb3/spotify"
)

// Start initializes complete program
func Start(username *string, pid *string) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(*pid)
	trackListJSON, _ := cli.UserClient.GetPlaylistTracks(*username, playlistID)
	for _, val := range trackListJSON.Tracks {
		fmt.Println("Added ", val.Track.Name)
		cli.TrackList = append(cli.TrackList, val.Track)
	}

	for _, val := range cli.TrackList {
		cli.YoutubeIDList = append(cli.YoutubeIDList, GetYoutubeIds(string(val.Name)+" "+string(val.Album.Name)+" music video"))
	}

	for _, track := range cli.YoutubeIDList {
		ytURL := "https://www.youtube.com/watch?v=" + track
		go Downloader(ytURL)
	}
}
