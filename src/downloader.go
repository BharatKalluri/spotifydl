package spotifydl

import "os/exec"

// Downloader is a function to download files
func Downloader(url string) {
	ytdlCmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3",
		"-o", "%(title)s.%(ext)s", "--audio-quality", "0", url)
	_, err := ytdlCmd.Output()
	if err != nil {
		panic(err)
	}
}
