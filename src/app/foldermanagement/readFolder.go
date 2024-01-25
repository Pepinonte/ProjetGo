package foldermanagement

import (
	"fmt"
	"os"
)

func ReadFolder(path string) {
	if path == "" {
		path = "."
	}
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
