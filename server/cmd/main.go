package main

import (
	"hara/internal/api"
	"hara/internal/convert"
)

func main() {
	converter := convert.NewConverter()
	converter.Initialize()

	server := api.NewServer("input", "output", converter)
	server.Run(":8080")
}
