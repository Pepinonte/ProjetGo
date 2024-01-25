package foldermanagement

import "os"

func DeleteFolder(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}