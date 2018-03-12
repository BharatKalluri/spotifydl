package spotifydl

import (
	"fmt"
	"os"

	"github.com/rylio/ytdl"
)

// Downloader is a function to download files
func Downloader(url string) {
	fmt.Println("Download started", url)

	vid, _ := ytdl.GetVideoInfo(url)
	audioTracks := ytdl.FormatList{}

	for _, format := range vid.Formats {
		if format.ValueForKey(ytdl.FormatResolutionKey) == "" {
			audioTracks = append(audioTracks, format)
		}
	}

	bestAudio := audioTracks.Best(ytdl.FormatAudioBitrateKey)
	if len(bestAudio) > 0 {
		file, err := os.Create(vid.Title + ".webm")
		if err != nil {
			panic("File Not Created: Check permissions")
		}
		defer file.Close()
		err = vid.Download(bestAudio[0], file)
		if err != nil {
			panic("Not Downloading, Check Internet")
		}
	}
}
