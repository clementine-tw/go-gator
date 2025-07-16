package main

import (
	"log"
	"os"

	"github.com/clementine-tw/go-gator/internal/config"
)

func main() {
	// read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	// initialize state
	curState := &state{
		cfg: &cfg,
	}

	// register commands
	registeredCommands := commands{
		Handlers: make(map[string]func(*state, command) error),
	}
	err = registeredCommands.register("login", handlerLogin)
	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	input := os.Args[1:]
	cmd := command{
		Name: input[0],
		Args: input[1:],
	}
	err = registeredCommands.run(curState, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}
}
