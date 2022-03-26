package utils

import (
	"fmt"
	"github.com/bogem/id3v2"
	"github.com/zmb3/spotify/v2"
	"log"
	"strconv"
	"strings"
	"time"
)

// TagFileWithSpotifyMetadata takes in a filename as a string and spotify metadata and uses it to tag the music
func TagFileWithSpotifyMetadata(fileName string, trackData spotify.FullTrack) {

	albumTag := trackData.Album.Name
	var trackArtist []string
	for _, Artist := range trackData.Album.Artists {
		trackArtist = append(trackArtist, Artist.Name)
	}
	artistTag := strings.Join(trackArtist[:], ",")
	dateObject, _ := time.Parse("2006-01-02", trackData.Album.ReleaseDate)
	yearTag := dateObject.Year()
	albumArtImages := trackData.Album.Images

	mp3File, err := id3v2.Open(fileName, id3v2.Options{Parse: true})
	if err != nil {
		panic(err)
	}
	defer func(mp3File *id3v2.Tag) {
		err := mp3File.Close()
		if err != nil {
			panic(err)
		}
	}(mp3File)

	mp3File.SetTitle(trackData.Name)
	mp3File.SetArtist(artistTag)
	mp3File.SetAlbum(albumTag)
	mp3File.SetYear(strconv.Itoa(yearTag))

	if len(albumArtImages) > 0 {
		albumArtURL := albumArtImages[0].URL
		albumArt, albumArtDownloadErr := DownloadFile(albumArtURL)
		if albumArtDownloadErr == nil {
			pic := id3v2.PictureFrame{
				Encoding:    id3v2.EncodingUTF8,
				MimeType:    "image/jpeg",
				PictureType: id3v2.PTFrontCover,
				Description: "Front cover",
				Picture:     albumArt,
			}
			mp3File.AddAttachedPicture(pic)
		} else {
			fmt.Println("An error occured while downloading album art ", err)
		}
	} else {
		fmt.Println("No album art found for ", trackData.Name)
	}

	if err = mp3File.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}

}
