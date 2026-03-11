package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Celudor/gogator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) >= 1 {
		cmdLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = cmdLimit
	} 

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return err
	}
	
	for _, post := range posts {
		fmt.Printf("%s\n%s\n%s\n\n", post.Title, post.Url, post.Description.String)
	}
	return nil

}
