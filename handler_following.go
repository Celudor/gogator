package main

import (
	"context"
	"fmt"

	"github.com/Celudor/gogator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("%s %s %s\n", feed.FeedName, feed.Url, feed.UserName)
	}
	return nil
}
