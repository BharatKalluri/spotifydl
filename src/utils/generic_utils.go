package utils

import (
	"io/ioutil"
	"net/http"
)

// DownloadFile will get a url and return bytes
func DownloadFile(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buffer, _ := ioutil.ReadAll(resp.Body)
	return buffer
}
