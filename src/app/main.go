package main

import (
	"app/interfaces"
	"os"
)

func main() {
	choiser(os.Args)
}

func choiser(arguements []string) {
	switch arguements[1] {
	case "createFolder", "deleteFolder", "readFolder", "renameFolder", "createFile", "deleteFile", "readFile", "modifyFile", "renameFile", "showLogs":
		interfaces.IcliModeMain(arguements)
	case "server":
		interfaces.IserverModeMain()
	}
}
