package rss_reader

import (
	"encoding/json"
	"fmt"
)

func parseToJson(feeds []*Feed) ([]*RssItem, error) {
	feed := Feed{}
	marsh, _ := json.Marshal(feeds)
	js := json.Unmarshal(marsh, &feed)

	if js != nil {
		return nil, fmt.Errorf("error :%w", js)
	}
	fmt.Println(feed)

	return nil, nil
}
