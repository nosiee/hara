package main

import (
	"hara/internal/api"
	"hara/internal/convert"
)

func main() {
	converter := convert.NewConverter("output")
	converter.Initialize()

	server := api.NewServer("input", converter)
	server.Run(":8080")
}
