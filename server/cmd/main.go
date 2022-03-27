package main

import (
	"hara/internal/api"
)

func main() {
	server := api.NewServer("input")
	server.Run(":8080")
}
