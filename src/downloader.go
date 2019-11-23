package spotifydl

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/BharatKalluri/spotifydl/src/utils"
	"github.com/bogem/id3v2"
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
	// TODO: Check images is not null
	albumArtURL := track.Album.Images[0].URL
	albumArt := utils.DownloadFile(albumArtURL)

	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", track.Name+".%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		panic(err)
	}

	// Tag the file with metadata
	mp3File, err := id3v2.Open(nameTag, id3v2.Options{Parse: true})
	if err != nil {
		panic(err)
	}
	defer mp3File.Close()
	mp3File.SetTitle(track.Name)
	mp3File.SetArtist(artistTag)
	mp3File.SetAlbum(albumTag)
	mp3File.SetYear(strconv.Itoa(yearTag))
	pic := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Description: "Front cover",
		Picture:     albumArt,
	}
	mp3File.AddAttachedPicture(pic)
	if err = mp3File.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}
}
