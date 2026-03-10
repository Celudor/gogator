package main

import "context"

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}
	return nil
}
