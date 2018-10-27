package spotifydl

import (
	"fmt"

	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	"github.com/zmb3/spotify"
)

// DownloadPlaylist Start initializes complete program
func DownloadPlaylist(pid string) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(pid)
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
	fmt.Println("👍 Found", len(cli.TrackList), "tracks")
	fmt.Println("🎵 Searching and downloading tracks")
	uiprogress.Start()
	bar := uiprogress.AddBar(len(cli.TrackList))

	bar.AppendCompleted()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		if b.Current() == len(cli.TrackList) {
			return "   🔍 " + strutil.Resize("Search complete", 30)
		}
		return "   🔍 " + strutil.Resize(cli.TrackList[b.Current()].Name, 30)
	})
	for _, val := range cli.TrackList {
		cli.YoutubeIDList = append(cli.YoutubeIDList, GetYoutubeIds(string(val.Name)+" "+string(val.Artists[0].Name)))
		bar.Incr()
	}
	bar2 := uiprogress.AddBar(len(cli.TrackList))
	bar2.AppendCompleted()
	bar2.PrependFunc(func(b *uiprogress.Bar) string {
		if b.Current() == len(cli.TrackList) {
			return "   ⬇️  " + strutil.Resize("Download complete", 30)
		}
		return "  ⬇️  " + strutil.Resize(fmt.Sprintf("Downloading: %s (%d/%d)", cli.TrackList[b.Current()].Name, b.Current(), len(cli.TrackList)), 30)
	})
	for index, track := range cli.YoutubeIDList {
		ytURL := "https://www.youtube.com/watch?v=" + track
		Downloader(ytURL, cli.TrackList[index])
		bar2.Incr()
	}
	uiprogress.Stop()
}
