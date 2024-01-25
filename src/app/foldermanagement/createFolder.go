package foldermanagement

import "os"

func CreateFolder(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		panic(err)
	}
}