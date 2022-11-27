package rss_reader

import (
	"encoding/json"
)

func parseToJson(feeds []*Feed) ([]*RssItem, error) {
	var rssArr []*RssItem
	for _, val := range feeds {
		itm := Feed{}
		item, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(item, &itm)
		if err != nil {
			return nil, err
		}

		for i, v := range itm.Entries {
			rssI := RssItem{
				v[i].Title,
				val.Title,
				val.Source[0].HREF,
				v[i].SourceURL[0].HREF,
				v[i].Published,
				v[i].Description,
			}

			rssArr = append(rssArr, &rssI)
		}

	}
	return rssArr, nil
}
