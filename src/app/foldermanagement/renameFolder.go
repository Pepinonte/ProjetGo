package foldermanagement

import "os"

func RenameFolder(path string, newPath string) error {
	return os.Rename(path, newPath)
}