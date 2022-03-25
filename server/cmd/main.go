package main

import (
	"hara/internal/api"
	"hara/internal/convert"
)

func main() {
	converter := convert.NewConverter()
	converter.Initialize()

	// TODO: I think it will be better, if we pass output path to converter??
	server := api.NewServer("input", "output", converter)
	server.Run(":8080")
}
