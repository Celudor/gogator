package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Celudor/gogator/internal/database"
	"github.com/Celudor/gogator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("time is required")
	}
	timeBetween, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetween)
	fmt.Printf("Collecting feeds every %v\n", timeBetween)
	for ;;<-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s* state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	if err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID: feed.ID,
	}); err != nil {
		return err
	}
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
