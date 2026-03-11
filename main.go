package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Celudor/gogator/internal/config"
	"github.com/Celudor/gogator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	s := state{}	
	if cfg, err := config.Read(); err != nil {
		fmt.Println(fmt.Errorf("Can't read config: %w", err))
		os.Exit(1)
	} else {
		s.cfg = &cfg
	}
	if db, err := sql.Open("postgres", s.cfg.DbUrl); err != nil {
		fmt.Println(fmt.Errorf("can't load db: %w", err))
		os.Exit(1)
	} else {
		s.db = database.New(db)
	}
	commands := commands{handlers: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))
	args := os.Args
	if len(args) < 2 {
		fmt.Println("command missing")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	err := commands.run(&s, command{name: cmdName, args: cmdArgs})
	if err != nil {
		fmt.Println(fmt.Errorf("error occurred: %w", err))
		os.Exit(1)
	}

}
