package main

import (
	"fmt"
	"os"

	"github.com/Celudor/gogator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(fmt.Errorf("Can't read config: %w", err))
		os.Exit(1)
	}
	s := state{cfg: &cfg}
	commands := commands{handlers: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		fmt.Println("command missing")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	err = commands.run(&s, command{name: cmdName, args: cmdArgs})
	if err != nil {
		fmt.Println(fmt.Errorf("error occurred: %w", err))
		os.Exit(1)
	}

}
