package rss_reader

import (
	"encoding/json"
)

func parseToJson(feeds []*Feed) ([]*RssItem, error) {
	var rssArr []*RssItem
	itm := Feed{}
	for _, val := range feeds {
		//itm := Feed{}
		item, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(item, &itm)
		if err != nil {
			return nil, err
		}

		for _, v := range itm.Entries {
			rssI := RssItem{
				v[0].Title,
				val.Title,
				val.Source[0].HREF,
				v[0].SourceURL[0].HREF,
				v[0].Published,
				v[0].Description,
			}

			rssArr = append(rssArr, &rssI)
		}

	}
	return rssArr, nil
}
