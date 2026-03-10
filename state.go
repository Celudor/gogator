package main

import (
	"github.com/Celudor/gogator/internal/config"
	"github.com/Celudor/gogator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}
