package main

import (
	"hara/internal/api"
	"hara/internal/convert"
)

func main() {
	// TODO: Separate folders for video input/output and image input/output
	converter := convert.NewConverter("output")
	converter.Initialize()

	server := api.NewServer("input", converter)
	server.Run(":8080")
}
