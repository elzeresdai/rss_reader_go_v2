package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"rss_reader_go_v2/pkg/rss_reader"
)

var urls struct {
	Urls []Url `json:"urls"`
}

type Url struct {
	Url string `json:"url"`
}

func GetRssData(e echo.Context) error {
	//var empty_slice []int
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		panic(err)
	}
	defer e.Request().Body.Close()

	err = json.Unmarshal(body, &urls)
	if err != nil {
		panic(err)
	}

	var feedUrls []string
	for _, val := range urls.Urls {
		url := val.Url
		feedUrls = append(feedUrls, url)
	}

	f, err := rss_reader.FetchFeeds(feedUrls)
	if err != nil {
		return err
	}
	fmt.Println(f)
	return nil
}
