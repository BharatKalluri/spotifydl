package spotifydl

import (
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var httpClient = &http.Client{}

type SearchResult struct {
	Title, Uploader, URL, Duration, ID string
	Live                               bool
	SourceName                         string
	Extra                              []string
}

// GetYoutubeId takes the query as string and returns the search results video ID's
func GetYoutubeId(searchQuery string) (string, error) {
	searchResults, err := ytSearch(searchQuery, 1)
	if err != nil {
		return "", err
	}
	if len(searchResults) == 0 {
		errorMessage := fmt.Sprintf("no songs found for %s", searchQuery)
		return "", errors.New(errorMessage)
	}
	return searchResults[0].ID, nil
}

func getContent(data []byte, index int) []byte {
	id := fmt.Sprintf("[%d]", index)
	contents, _, _, _ := jsonparser.Get(data, "contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer", "contents", id, "itemSectionRenderer", "contents")
	return contents
}

// shamelessly ripped off from https://github.com/Pauloo27/tuner/blob/11dd4c37862c1c26521a01c8345c22c29ab12749/search/youtube.go#L27

func ytSearch(searchTerm string, limit int) (results []*SearchResult, err error) {
	ytSearchUrl := fmt.Sprintf(
		"https://www.youtube.com/results?search_query=%s", url.QueryEscape(searchTerm),
	)

	req, err := http.NewRequest("GET", ytSearchUrl, nil)
	if err != nil {
		return nil, errors.New("cannot get youtube page")
	}
	req.Header.Add("Accept-Language", "en")
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, errors.New("cannot get youtube page")
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != 200 {
		return nil, errors.New("failed to make a request to youtube")
	}

	buffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("cannot read response from youtube")
	}

	body := string(buffer)
	splitScript := strings.Split(body, `window["ytInitialData"] = `)
	if len(splitScript) != 2 {
		splitScript = strings.Split(body, `var ytInitialData = `)
	}

	if len(splitScript) != 2 {
		return nil, errors.New("invalid response from youtube")
	}
	splitScript = strings.Split(splitScript[1], `window["ytInitialPlayerResponse"] = null;`)
	jsonData := []byte(splitScript[0])

	index := 0
	var contents []byte

	for {
		contents = getContent(jsonData, index)
		_, _, _, err = jsonparser.Get(contents, "[0]", "carouselAdRenderer")

		if err == nil {
			index++
		} else {
			break
		}
	}

	_, err = jsonparser.ArrayEach(contents, func(value []byte, t jsonparser.ValueType, i int, err error) {
		if err != nil {
			return
		}

		if limit > 0 && len(results) >= limit {
			return
		}

		id, err := jsonparser.GetString(value, "videoRenderer", "videoId")
		if err != nil {
			return
		}

		title, err := jsonparser.GetString(value, "videoRenderer", "title", "runs", "[0]", "text")
		if err != nil {
			return
		}

		uploader, err := jsonparser.GetString(value, "videoRenderer", "ownerText", "runs", "[0]", "text")
		if err != nil {
			return
		}

		live := false
		duration, err := jsonparser.GetString(value, "videoRenderer", "lengthText", "simpleText")

		if err != nil {
			duration = ""
			live = true
		}

		results = append(results, &SearchResult{
			Title:      title,
			Uploader:   uploader,
			Duration:   duration,
			ID:         id,
			URL:        fmt.Sprintf("https://youtube.com/watch?v=%s", id),
			Live:       live,
			SourceName: "youtube",
		})
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
