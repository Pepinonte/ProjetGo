package foldermanagement

import (
	"fmt"
	"os"
)

func ReadFolder(path string) error{
	if path == "" {
		path = "."
	}
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}