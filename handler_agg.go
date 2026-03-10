package main

import (
	"context"
	"fmt"

	"github.com/Celudor/gogator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}
	fmt.Println(rssFeed)
	return nil
}
