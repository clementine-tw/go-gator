package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/clementine-tw/go-gator/internal/config"
	"github.com/clementine-tw/go-gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	// connect db
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting database: %v", err)
	}

	dbQueries := database.New(db)
	// initialize state
	curState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	// register commands
	toRegCmds := map[string]func(*state, command) error{
		"login":     handlerLogin,
		"register":  handlerRegister,
		"reset":     handlerReset,
		"users":     handlerUsers,
		"agg":       handlerAggregate,
		"addfeed":   middlewareLoggedIn(handlerAddFeed),
		"feeds":     handlerFeeds,
		"follow":    middlewareLoggedIn(handlerFollowFeed),
		"following": middlewareLoggedIn(handlerListFeedFollows),
		"unfollow":  middlewareLoggedIn(handlerUnfollowFeed),
		"browse":    middlewareLoggedIn(handlerBrowse),
	}
	cmds := commands{
		Handlers: make(map[string]func(*state, command) error),
	}
	for name, handler := range toRegCmds {

		err = cmds.register(name, handler)
		if err != nil {
			log.Fatalf("error registering command: %v", err)
		}
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	input := os.Args[1:]
	cmd := command{
		Name: input[0],
		Args: input[1:],
	}
	err = cmds.run(curState, cmd)
	if err != nil {
		log.Fatalf("error running command '%v': %v", cmd.Name, err)
	}
}
