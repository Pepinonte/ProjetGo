package main

import (
	"app/filesmanagement"
	"app/foldermanagement"
	"os"
)

func main() {

	switch os.Args[1] {
	case "createFolder":
		foldermanagement.CreateFolder(os.Args[2])
	case "readFolder":
		if len(os.Args) >= 3 {
			foldermanagement.ReadFolder(os.Args[2])
			
		} else {
			foldermanagement.ReadFolder("")
		}
	case "deleteFolder":
		foldermanagement.DeleteFolder(os.Args[2])
	case "renameFolder":
		foldermanagement.RenameFolder(os.Args[2], os.Args[3])
	case "createFile":
		filesmanagement.CreateFile(os.Args[2])
	case "deleteFile":
			filesmanagement.DeleteFile(os.Args[2])		
	case "readFile":
			filesmanagement.ReadFile(os.Args[2])
	case "modifyFile":
		filesmanagement.ModifyFile(os.Args[2], os.Args[3])
	case "renameFile":
		filesmanagement.RenameFile(os.Args[2], os.Args[3])
	}

	


}
