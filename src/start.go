package spotifydl

import (
	"fmt"

	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	"github.com/zmb3/spotify"
)

// Start initializes complete program
func Start(username *string, pid *string) {
	user := InitAuth()
	cli := UserData{
		UserClient: user,
	}
	playlistID := spotify.ID(*pid)
	trackListJSON, _ := cli.UserClient.GetPlaylistTracks(playlistID)
	for _, val := range trackListJSON.Tracks {
		cli.TrackList = append(cli.TrackList, val.Track)
	}

	fmt.Println("👍 Found", len(cli.TrackList), "tracks")
	fmt.Println("🎵 Searching and downloading tracks")
	uiprogress.Start()
	bar := uiprogress.AddBar(len(cli.TrackList))

	bar.AppendCompleted()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		if b.Current() == len(cli.TrackList){
			return "   🔍 " + strutil.Resize("Search complete", 30)
		}
		return "   🔍 " + strutil.Resize(cli.TrackList[b.Current()].Name, 30)
	})
	for _, val := range cli.TrackList {
		cli.YoutubeIDList = append(cli.YoutubeIDList, GetYoutubeIds(string(val.Name)+" "+string(val.Album.Name)+" music video"))
		bar.Incr()
	}
	bar2 := uiprogress.AddBar(len(cli.TrackList))
	bar2.AppendCompleted()
	bar2.PrependFunc(func(b *uiprogress.Bar) string {
		if b.Current() == len(cli.TrackList){
			return "   ⬇️  " + strutil.Resize("Download complete", 30)
		}
		return "   ⬇️  " + strutil.Resize(cli.TrackList[b.Current()].Name, 30)
	})
	for _, track := range cli.YoutubeIDList {
		ytURL := "https://www.youtube.com/watch?v=" + track
		Downloader(ytURL)
		bar2.Incr()
	}
	uiprogress.Stop()
}
