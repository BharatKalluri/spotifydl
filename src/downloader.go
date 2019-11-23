package spotifydl

import (
	"fmt"
	"os/exec"

	"github.com/zmb3/spotify"

	"github.com/BharatKalluri/spotifydl/src/utils"
)

// Downloader is a function to download files
func Downloader(url string, track spotify.FullTrack) {
	nameTag := fmt.Sprintf("%s.mp3", track.Name)

	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", track.Name+".%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		panic(err)
	}

	// Tag the file with metadata
	utils.TagFileWithSpotifyMetadata(nameTag, track)
}
