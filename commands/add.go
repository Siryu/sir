package commands

import (
	"crypto/sha1"

	"github.com/Siryu/sir/data"
)

func Add(files []data.File) (err error) {

	for _, file := range files {
		go func(file string) {
			newNode := data.Node{
				Name: file.Name,
				Hash: sha1.Sum([]byte(file.Data)),
			}
		}(file)
	}
}
