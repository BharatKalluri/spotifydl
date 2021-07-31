package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// DownloadFile will get a url and return bytes
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		HandleError(err, fmt.Sprintf("Could not download from URL: %s", url))
	}(resp.Body)
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to download album art!")
		return nil, err
	}
	return buffer, nil
}

func HandleError(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}
