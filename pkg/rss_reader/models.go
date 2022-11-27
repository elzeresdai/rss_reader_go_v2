package rss_reader

import "time"

type Items struct {
	RssItem []RssItem `json:"RssItem"`
}

type RssItem struct {
	Title       string    `json:"title"`
	Source      string    `json:"source"`
	SourceURL   string    `json:"source_url"`
	Links       string    `json:"links"`
	PublishDate time.Time `json:"publish_date"`
	Description string    `json:"description"`
}

type Feed struct {
	Title   string  `xml:"title" json:",string"`
	Source  []Link  `xml:"link,href" json:",string"`
	Entries []Entry `xml:"entry" json:",string"`
}

type Entry []struct {
	Title       string    `xml:"title" json:",string"`
	SourceURL   []Link    `xml:"link,href" json:",string"`
	Description string    `xml:"summary" json:",string"`
	Published   time.Time `xml:"published" json:",string"`
}

type Link struct {
	HREF string `xml:"href,attr" json:",string"`
}
