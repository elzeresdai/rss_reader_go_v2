package rss_reader

import "time"

type RssItem struct {
	Title       string `json:"Title"`
	Source      string `json:"Source"`
	SourceURL   string `json:"SourceURL"`
	Links       []Link `json:"Link"`
	PublishDate time.Time
	Description string `json:"Description"`
}

type Feed struct {
	Title   string  `xml:"title" json:"title"`
	Link    []Link  `xml:"link,href" json:"link"`
	Entries []Entry `xml:"entry" json:"entries"`
}

type Entry struct {
	Title     string    `xml:"title" json:"title"`
	Links     []Link    `xml:"link,href" json:"links"`
	Summary   string    `xml:"summary" json:"summary"`
	Published time.Time `xml:"published" json:"published"`
}

type Link struct {
	HREF string `xml:"href,attr" json:"HREF"`
}
