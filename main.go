package main

import (
	"fmt"
	"os"

	spotifydl "github.com/BharatKalluri/spotifydl/src"
	"github.com/spf13/cobra"
)

func main() {
	var playlistid string
	var albumid string

	var rootCmd = &cobra.Command{
		Use:   "spotifydl",
		Short: "spotifydl is a awesome music downloader",
		Long: `Spotifydl lets you download albums and playlists and tags them for you
Pass Either album ID or Playlist ID to start downloading`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(playlistid) > 0 && len(albumid) > 0 {
				fmt.Println("Either album ID or playlist ID")
				cmd.Help()
			} else if len(albumid) > 0 {
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
