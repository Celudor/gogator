package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Celudor/gogator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("Name and url is required")
	}


	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}
	if err := handlerFollow(s, command{name: "", args: []string{cmd.args[1]}}, user); err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
