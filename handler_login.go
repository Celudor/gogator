package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("command require username")
	}

	
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		return errors.New("username doesn't exist")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set\n", cmd.args[0])
	return nil
}
