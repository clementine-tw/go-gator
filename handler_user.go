package main

import (
	"context"
	"fmt"
	"time"

	"github.com/clementine-tw/go-gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, c command) error {

	if len(c.Args) == 0 {
		return fmt.Errorf("usage: %s <user_name>", c.Name)
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
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("ID: %s\n", user.ID)
	return nil
}

func handlerLogin(s *state, c command) error {

	if len(c.Args) == 0 {
		return fmt.Errorf("usage: %s <user_name>", c.Name)
	}

	username := c.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Login successfully, current user: %s\n", username)

	return nil
}

func handlerReset(s *state, _ command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset users: %w", err)
	}

	return nil
}
