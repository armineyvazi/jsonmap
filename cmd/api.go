package main

import (
	"github.com/armineyvazi/jsonmap/di"
	_ "github.com/armineyvazi/jsonmap/docs"
	"os"
)

// @title JSONMap API
// @version 1.0
// @description API for interacting with the JSONMap GPT service.
// @host localhost:3000
// @BasePath /api/v1
func main() {
	if err := di.InitializeApp(); err != nil {
		os.Exit(1)
	}
}
