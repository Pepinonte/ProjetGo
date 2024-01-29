package main

import (
	"app/interfaces"
	"fmt"
	"os"
)

const conMode = "online"

func main() {
	choiser(conMode, os.Args)
}

func choiser(conMode string, arguements []string) {
	switch arguements[1] {
	case "createFolder", "deleteFolder", "readFolder", "renameFolder", "createFile", "deleteFile", "readFile", "modifyFile", "renameFile", "showLogs":
		interfaces.IcliModeMain(conMode, arguements)
	case "server":
		interfaces.IserverModeMain()
	default:
		fmt.Println("Invalid command")
	}
}
