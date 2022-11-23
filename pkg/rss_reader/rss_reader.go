package rss_reader

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type feedResult struct {
	feed *Feed
	err  error
}

func FetchFeeds(urls []string) []*Feed {
	feedCh := make(chan feedResult)
	for _, url := range urls {
		go fetchUrl(url, feedCh)
	}

	var feeds []*Feed
	for i := 0; i < len(urls); i++ {
		res := <-feedCh
		// If the goroutine errors out, we'll just wait for others
		if res.err != nil {
			continue
		}
		feeds = append(feeds, res.feed)
	}

	return feeds

}

func fetchUrl(url string, feedCh chan feedResult) {
	// Create a client with a default timeout
	net := &http.Client{
		Timeout: time.Second * 10,
	}

	// get data from url and check error
	res, err := net.Get(url)
	if err != nil {
		feedCh <- feedResult{nil, err}
		return
	}

	defer res.Body.Close()

	// Read the body of the request and parse the feed
	body, err := io.ReadAll(res.Body)
	if err != nil {
		feedCh <- feedResult{nil, err}
	}

	feed, err := parseBody(body)
	if err != nil {
		feedCh <- feedResult{nil, err}
		return
	}
	feedCh <- feedResult{feed, nil}

}

func parseBody(body []byte) (*Feed, error) {
	feed := Feed{}
	err := xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}
	return &feed, nil
}
