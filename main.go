package main

import (
	"log"

	"github.com/Siryu/sir/commands"
	"github.com/Siryu/sir/data"
)

func main() {
	// let's start with add
	command := "add"
	file := []data.File{
		data.File{
			name: "pie",
			data: []byte("apple pie"),
		},
	}
	head := new(data.Head)

	switch command {
	case "add":
		commands.Add(file)
	case "commit":
	case "checkout":
	default:
		log.Println("Select command -- add -- commit -- checkout")
	}
}
