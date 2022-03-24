package main

import "hara/internal/api"

func main() {
	server := api.NewServer("input/", "output/")
	server.Run(":8080")
}
