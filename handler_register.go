package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Celudor/gogator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister (s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("username is missing")
	}

	usr, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: cmd.args[0],
		})
	if err != nil {
		return err
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Println("user was crated")
	fmt.Printf("ID: %v, Created At:%v, Update At: %v, Name: %v\n", usr.ID, usr.CreatedAt, usr.UpdatedAt, usr.Name)
	return nil

}  
