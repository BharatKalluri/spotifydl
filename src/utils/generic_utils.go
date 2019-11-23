package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// DownloadFile will get a url and return bytes
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to download album art!")
		return nil, err
	}
	return buffer, nil
}
