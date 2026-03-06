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
	cfg.SetUser("lukasz")
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(fmt.Errorf("Can't read config: %w", err))
		os.Exit(1)
	}

	fmt.Println(cfg.DbUrl)
	fmt.Println(cfg.CurrentUserName)
}
