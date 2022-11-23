package rss_reader

import (
	"time"
)

type RssItem struct {
	Title       string `json:"Title"`
	Source      string `json:"Source"`
	SourceURL   string `json:"SourceURL"`
	Link        string `json:"Link"`
	PublishDate time.Time
	Description string `json:"Description"`
}

type Feed struct {
	Title   string  `xml:"title"`
	Links   []Link  `xml:"link"`
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Title     string `xml:"title"`
	Links     []Link `xml:"link"`
	Summary   string `xml:"summary"`
	Published string `xml:"published"`
}

type Link struct {
	HREF string `xml:"href,attr"`
}
