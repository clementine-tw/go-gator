package main

import (
	"fmt"
	"os"

	"github.com/clementine-tw/go-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		os.Exit(1)
	}

	err = cfg.SetUser("Clement")
	if err != nil {
		fmt.Printf("error setting current user name: %v\n", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("config file: %v\n", cfg)
}
