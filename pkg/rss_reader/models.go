package rss_reader

import "time"

type Items struct {
	RssItem []RssItem `json:"RssItem"`
}

type RssItem struct {
	Title       string    `json:"title,omitempty"`
	Source      string    `json:"source,omitempty"`
	SourceURL   string    `json:"source_url,omitempty"`
	Links       string    `json:"links,omitempty"`
	PublishDate time.Time `json:"publish_date,omitempty"`
	Description string    `json:"description,omitempty"`
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
