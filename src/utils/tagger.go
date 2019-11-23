package utils

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/bogem/id3v2"
	"github.com/zmb3/spotify"
)

// TagFileWithSpotifyMetadata takes in a filename as a string and spotify metadata and uses it to tag the music
func TagFileWithSpotifyMetadata(fileName string, trackData spotify.FullTrack) {

	albumTag := trackData.Album.Name
	trackArtist := []string{}
	for _, Artist := range trackData.Album.Artists {
		trackArtist = append(trackArtist, Artist.Name)
	}
	artistTag := strings.Join(trackArtist[:], ",")
	dateObject, _ := time.Parse("2006-01-02", trackData.Album.ReleaseDate)
	yearTag := dateObject.Year()
	// TODO: Check images is not null
	albumArtURL := trackData.Album.Images[0].URL
	albumArt := DownloadFile(albumArtURL)

	mp3File, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		panic(err)
	}
	defer mp3File.Close()

	mp3File.SetTitle(trackData.Name)
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
