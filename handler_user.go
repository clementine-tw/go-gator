package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, c command) error {

	if len(c.Args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	username := c.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting user name when login: %w", err)
	}

	fmt.Printf("%s has been set\n", username)

	return nil
}
