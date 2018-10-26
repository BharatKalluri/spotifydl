package spotifydl

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	id3 "github.com/mikkyang/id3-go"
	"github.com/zmb3/spotify"
)

// Downloader is a function to download files
func Downloader(url string, track spotify.FullTrack) {
	nameTag := fmt.Sprintf("%s.mp3", track.Name)
	albumTag := track.Album.Name
	trackArtist := []string{}
	for _, Artist := range track.Album.Artists {
		trackArtist = append(trackArtist, Artist.Name)
	}
	artistTag := strings.Join(trackArtist[:], ",")
	dateObject, _ := time.Parse("2006-01-02", track.Album.ReleaseDate)
	yearTag := dateObject.Year()

	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", track.Name+".%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		panic(err)
	}

	// Tag the file with metadata
	mp3File, err := id3.Open(nameTag)
	if err != nil {
		panic(err)
	}
	defer mp3File.Close()
	mp3File.SetTitle(track.Name)
	mp3File.SetArtist(artistTag)
	mp3File.SetAlbum(albumTag)
	mp3File.SetYear(strconv.Itoa(yearTag))
}
