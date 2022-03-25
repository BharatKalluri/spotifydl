package spotifydl

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"log"
	"os"
	"strings"
)

// DownloadPlaylist Start initializes complete program
func DownloadPlaylist(pid string, ctx context.Context) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(pid)

	trackListJSON, err := cli.UserClient.GetPlaylistTracks(ctx, playlistID)
	if err != nil {
		fmt.Println("Playlist not found!")
		os.Exit(1)
	}
	for _, val := range trackListJSON.Tracks {
		cli.TrackList = append(cli.TrackList, val.Track)
	}

	for page := 0; ; page++ {
		err := cli.UserClient.NextPage(ctx, trackListJSON)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, val := range trackListJSON.Tracks {
			cli.TrackList = append(cli.TrackList, val.Track)
		}
	}

	DownloadTrackList(cli)
}

// DownloadAlbum Download album according to
func DownloadAlbum(aid string, ctx context.Context) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	albumID := spotify.ID(aid)
	album, err := user.GetAlbum(ctx, albumID)
	if err != nil {
		fmt.Println("Album not found!")
		os.Exit(1)
	}
	for _, val := range album.Tracks.Tracks {
		cli.TrackList = append(cli.TrackList, spotify.FullTrack{
			SimpleTrack: val,
			Album:       album.SimpleAlbum,
		})
	}
	DownloadTrackList(cli)
}

// DownloadSong will download a song with its identifier
func DownloadSong(sid string, ctx context.Context) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	songID := spotify.ID(sid)
	song, err := user.GetTrack(ctx, songID)
	if err != nil {
		fmt.Println("Song not found!")
		os.Exit(1)
	}
	cli.TrackList = append(cli.TrackList, spotify.FullTrack{
		SimpleTrack: song.SimpleTrack,
		Album:       song.Album,
	})
	DownloadTrackList(cli)
}

// DownloadTrackList Start downloading given list of tracks
func DownloadTrackList(cli UserData) {
	fmt.Println("Found", len(cli.TrackList), "tracks")
	fmt.Println("Searching and downloading tracks")
	for _, val := range cli.TrackList {
		var artistNames []string
		for _, artistInfo := range val.Artists {
			artistNames = append(artistNames, artistInfo.Name)
		}
		searchTerm := strings.Join(artistNames, " ") + " " + val.Name
		youtubeID, err := GetYoutubeId(searchTerm, val.Duration/1000)
		if err != nil {
			log.Printf("Error occured for %s error: %s", val.Name, err)
			continue
		}
		cli.YoutubeIDList = append(cli.YoutubeIDList, youtubeID)
	}
	for index, track := range cli.YoutubeIDList {
		fmt.Println()
		ytURL := "https://www.youtube.com/watch?v=" + track
		fmt.Println("⇓ Downloading " + cli.TrackList[index].Name)
		Downloader(ytURL, cli.TrackList[index].SimpleTrack)
		fmt.Println()
	}
	fmt.Println("Download complete!")
}
