package main

import (
	"app/interfaces"
	"os"
)

func main() {

	choiser(os.Args)
	// interfaces.IserverModeMain()

	// database.Condb()
	// var arguments string
	// currentTime := time.Now()

	// switch os.Args[1] {
	// case "createFolder":
	// 	outputCaseCreateFolder, err := foldermanagement.CreateFolder(os.Args[2])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "CreateFolder()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "CreateFolder()", arguments, outputCaseCreateFolder)
	// 	}
	// case "readFolder":
	// 	if len(os.Args) >= 3 {
	// 		outputCaseReadFolder, err := foldermanagement.ReadFolder(os.Args[2])
	// 		for i := 1; i < len(os.Args); i++ {
	// 			arguments += os.Args[i] + " "
	// 		}

	// 		if err != nil {
	// 			database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", arguments, err.Error())
	// 		} else {
	// 			database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", arguments, outputCaseReadFolder)
	// 		}
	// 	} else {
	// 		outputCaseReadFolder, err := foldermanagement.ReadFolder("")
	// 		for i := 1; i < len(os.Args); i++ {
	// 			arguments += os.Args[i] + " "
	// 		}

	// 		if err != nil {
	// 			database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", arguments, err.Error())
	// 		} else {
	// 			database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "ReadFolder()", arguments, outputCaseReadFolder)
	// 		}
	// 	}
	// case "deleteFolder":

	// 	outputCaseDeleteFolder, err := foldermanagement.DeleteFolder(os.Args[2])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "DeleteFolder()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "DeleteFolder()", arguments, outputCaseDeleteFolder)
	// 	}
	// case "renameFolder":

	// 	outputCaseRenameFolder, err := foldermanagement.RenameFolder(os.Args[2], os.Args[3])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "RenameFolder()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "foldermanagement", "RenameFolder()", arguments, outputCaseRenameFolder)
	// 	}

	// 	foldermanagement.RenameFolder(os.Args[2], os.Args[3])
	// case "createFile":

	// 	outputCaseCreateFile, err :=  filesmanagement.CreateFile(os.Args[2])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "CreateFile()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filesmanagement", "CreateFile()", arguments, outputCaseCreateFile)
	// 	}
	// case "deleteFile":

	// 	outputCaseDeleteFile, err := filesmanagement.DeleteFile(os.Args[2])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "DeleteFile()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "DeleteFile()", arguments, outputCaseDeleteFile)
	// 	}
	// case "readFile":

	// 	outputCaseReadFile, err := filesmanagement.ReadFile(os.Args[2])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "ReadFile()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "ReadFile()", arguments, outputCaseReadFile)
	// 	}
	// case "modifyFile":

	// 	outputCaseModifyFile, err := filesmanagement.ModifyFile(os.Args[2], os.Args[3])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "ModifyFile()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "ModifyFile()", arguments, outputCaseModifyFile)
	// 	}
	// case "renameFile":

	// 	outputCaseRenameFile, err := filesmanagement.RenameFile(os.Args[2], os.Args[3])
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}

	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filemanagement", "RenameFile()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "filenagement", "RenameFile()", arguments, outputCaseRenameFile)
	// 	}
	// case "showLogs":

	// 	outputCaseShowLog,  err := database.GetLogs()
	// 	for i := 1; i < len(os.Args); i++ {
	// 		arguments += os.Args[i] + " "
	// 	}
	// 	if err != nil {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "database", "GetLogs()", arguments, err.Error())
	// 	} else {
	// 		database.AddLog(currentTime.Format("2006.01.02 15:04:05"), "database", "GetLogs()", arguments, outputCaseShowLog)
	// 	}
	// }
}

// func choiser(arguements []string) {
// 	switch arguements[1] {
// 	case "createFolder", "deleteFolder", "readFolder", "renameFolder":
// 		interfaces.IfolderMain(arguements)
// 	// case "createFile", "deleteFile", "readFile", "modifyFile", "renameFile":
// 	// 	interfaces.IfileMain()
// 	}
// }

func choiser(arguements []string) {
	switch arguements[1] {
	case "createFolder", "deleteFolder", "readFolder", "renameFolder", "createFile", "deleteFile", "readFile", "modifyFile", "renameFile", "showLogs":
		interfaces.IcliModeMain(arguements)
	case "server":
		interfaces.IserverModeMain()
	}
}