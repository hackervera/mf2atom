package mf2atom

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/feeds"
	"willnorris.com/go/microformats"
)

func Parse(_url string) string {
	resp, err := http.Get(_url)
	if err != nil {
		log.Panic(err)
	}
	urlparsed, _ := url.Parse(_url)
	data := microformats.Parse(resp.Body, urlparsed)

	feed := &feeds.Feed{
		Title: _url,
		Link:  &feeds.Link{Href: _url},
	}

	items := []*feeds.Item{}
	for _, item := range data.Items {
		if item.Type[0] == "h-entry" {
			created := parseTime(item)
			new_item := &feeds.Item{
				Title:   item.Properties["name"][0].(string),
				Link:    &feeds.Link{Href: item.Properties["url"][0].(string)},
				Created: created,
			}
			items = append(items, new_item)
		}
	}
	feed.Items = items
	response, err := feed.ToAtom()
	if err != nil {
		log.Panic(err)
	}
	return response
}

func parseTime(item *microformats.Microformat) time.Time {
	created, _ := time.Parse(time.RFC3339, item.Properties["published"][0].(string))
	return created
}
