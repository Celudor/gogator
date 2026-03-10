package main

import (
	"context"
	"errors"

	"github.com/Celudor/gogator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("url is required")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.db.Unfollow(context.Background(), database.UnfollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
