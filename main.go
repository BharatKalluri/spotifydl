package main

import (
	"flag"
    "github.com/BharatKalluri/spotifydl/src"
)

func main() {
	username := flag.String("username", "Spotify", "Username of Playlist owner")
	pid := flag.String("playlistid", "37i9dQZF1DXcBWIGoYBM5M", "Playlist ID")

	flag.Parse()

	spotifydl.Start(username, pid)
}
