package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	Handlers map[string]func(s *state, c command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found: %s", cmd.Name)
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, handler func(*state, command) error) error {

	if _, exists := c.Handlers[name]; exists {
		return fmt.Errorf("error registering duplicate command: %s", name)
	}

	c.Handlers[name] = handler

	return nil
}
