package spotifydl

import (
	"fmt"

	"github.com/zmb3/spotify"
)

// DownloadPlaylist Start initializes complete program
func DownloadPlaylist(pid string) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(pid)
	// TODO: Exit gracefully if the playlist is not found
	trackListJSON, _ := cli.UserClient.GetPlaylistTracks(playlistID)
	for _, val := range trackListJSON.Tracks {
		cli.TrackList = append(cli.TrackList, val.Track)
	}
	DownloadTracklist(cli)
}

// DownloadAlbum Download album according to
func DownloadAlbum(aid string) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	albumid := spotify.ID(aid)
	// TODO: Exit gracefully if album is not found
	album, _ := user.GetAlbum(albumid)
	for _, val := range album.Tracks.Tracks {
		cli.TrackList = append(cli.TrackList, spotify.FullTrack{
			SimpleTrack: val,
			Album:       album.SimpleAlbum,
		})
	}
	DownloadTracklist(cli)
}

// DownloadTracklist Start downloading given list of tracks
func DownloadTracklist(cli UserData) {
	fmt.Println("Found", len(cli.TrackList), "tracks")
	fmt.Println("Searching and downloading tracks")
	for _, val := range cli.TrackList {
		cli.YoutubeIDList = append(cli.YoutubeIDList, GetYoutubeIds(string(val.Name)+" "+string(val.Artists[0].Name)))
	}
	for index, track := range cli.YoutubeIDList {
		fmt.Println()
		ytURL := "https://www.youtube.com/watch?v=" + track
		fmt.Println("⇓ Downloading " + cli.TrackList[index].Name)
		Downloader(ytURL, cli.TrackList[index])
		fmt.Println()
	}
	fmt.Println("Download complete!")
}
