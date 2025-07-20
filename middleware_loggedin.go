package main

import (
	"context"
	"fmt"

	"github.com/clementine-tw/go-gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, c command, user database.User) error) func(s *state, c command) error {

	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}

		return handler(s, c, user)
	}
}
