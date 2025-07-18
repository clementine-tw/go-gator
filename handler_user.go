package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/clementine-tw/go-gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, c command) error {

	if len(c.Args) == 0 {
		return errors.New("the register handler expects a single argument, the username")
	}
	username := c.Args[0]

	now := time.Now()
	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Name:      username,
		})

	if err != nil {
		return fmt.Errorf("error inserting user to db: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	log.Printf("successfully register user: %v", user)
	return nil
}

func handlerLogin(s *state, c command) error {

	if len(c.Args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	username := c.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user name when login: %w", err)
	}

	fmt.Printf("%s has been set\n", username)

	return nil
}
