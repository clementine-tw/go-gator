package main

import (
	"github.com/clementine-tw/go-gator/internal/config"
	"github.com/clementine-tw/go-gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
