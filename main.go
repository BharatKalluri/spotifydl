package main

import (
	"fmt"
	"os"
	"strings"

	spotifydl "github.com/BharatKalluri/spotifydl/src"
	"github.com/spf13/cobra"
)

func main() {
	var playlistid string
	var albumid string
	var spotifyURL string

	var rootCmd = &cobra.Command{
		Use:   "spotifydl",
		Short: "spotifydl is a awesome music downloader",
		Long: `Spotifydl lets you download albums and playlists and tags them for you.
Pass Either album ID or Playlist ID to start downloading`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(spotifyURL) > 0 {
				splitURL := strings.Split(spotifyURL, "/")

				if len(splitURL) < 2 {
					fmt.Println("Please enter the url copied from the spotify client")
					os.Exit(1)
				}

				spotifyID := splitURL[len(splitURL)-1]
				if strings.Contains(spotifyID, "?") {
					fmt.Println("Please remove the part of the url after the question mark (?) and try again")
					fmt.Println("For example, https://open.spotify.com/playlist/randomID?si=otherRandomID should just be https://open.spotify.com/playlist/randomID ")
					os.Exit(1)
				}

				if strings.Contains(spotifyURL, "album") {
					albumid = spotifyID
				} else if strings.Contains(spotifyURL, "playlist") {
					playlistid = spotifyID
				}
			}

			if len(albumid) > 0 {
				// Download album with the given album ID
				spotifydl.DownloadAlbum(albumid)
			} else if len(playlistid) > 0 {
				// Download playlist with the given ID
				spotifydl.DownloadPlaylist(playlistid)
			} else {
				fmt.Println("Enter valid input.")
				cmd.Help()
			}
		},
	}

	rootCmd.Flags().StringVarP(&playlistid, "playlistid", "p", "", "Album ID found on spotify")
	rootCmd.Flags().StringVarP(&albumid, "albumid", "a", "", "Album ID found on spotify")
	rootCmd.Flags().StringVarP(&spotifyURL, "spotifyurl", "u", "", "URL copied on spotify")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
