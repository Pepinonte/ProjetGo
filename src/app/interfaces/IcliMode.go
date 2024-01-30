package interfaces

import (
	"app/database"
	"app/filesmanagement"
	"app/foldermanagement"
)

type IcliMode interface {
	CreateFolder() (string, error)
	DeleteFolder() (string, error)
	ReadFolder() (string, error)
	RenameFolder() (string, error)
	CreateFile() (string, error)
	DeleteFile() (string, error)
	ReadFile() (string, error)
	RenameFile() (string, error)
	ShowLogs() (string, error)
}

// TODO: try no define the struct here and use the one from app/structures/Schoise.go
type Schoise struct {
	mode      string
	arguments []string
	conMode	 string
}

type Sfile struct {
	ID    		string `json:"id"`
	Name		string `json:"name"`
	Content	string `json:"content"`
}

func IcliModeMain(conMode string,arguments []string) {
	mycli := Schoise{arguments[1], arguments, conMode}
	execute(&mycli ,mycli)
}

// TODO: choise where set the switch logic (in the individual interfaces files or in a main file?)

func execute(cl IcliMode, st Schoise) {
	switch st.mode {
	case "createFolder":
		cl.CreateFolder()
	case "deleteFolder":
		cl.DeleteFolder()
	case "readFolder":
		cl.ReadFolder()
	case "renameFolder":
		cl.RenameFolder()
	case "createFile":
		cl.CreateFile()
	case "deleteFile":
		cl.DeleteFile()
	case "readFile":
		cl.ReadFile()
	case "renameFile":
		cl.RenameFile()
	case "showLogs":
		cl.ShowLogs()	
	}
}

func (cl *Schoise) CreateFolder() (string, error) {
	if cl.conMode == "offline" {
		foldermanagement.CreateFolder(cl.arguments[2])
	} else if cl.conMode == "online" {
		foldermanagement.ReqCreateFolder(cl.arguments[2])
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) DeleteFolder() (string, error) {
	if cl.conMode == "offline" {
		foldermanagement.DeleteFolder(cl.arguments[2])
	} else if cl.conMode == "online" {
		foldermanagement.ReqDeleteFolder(cl.arguments[2])
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) ReadFolder() (string, error) {
	if cl.conMode == "offline" {
		foldermanagement.ReadFolder(cl.arguments[2])
	} else if cl.conMode == "online" {
		foldermanagement.ReqReadFolder(cl.arguments[2])
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) RenameFolder() (string, error) {
	if cl.conMode == "offline" {
		foldermanagement.RenameFolder(cl.arguments[2], cl.arguments[3])
	} else if cl.conMode == "online" {
		// foldermanagement.ReqRenameFolder(cl.arguments[2], cl.arguments[3])
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) CreateFile() (string, error) {
	myFile, _ := filesmanagement.CreateFile(cl.arguments[2])
	if cl.conMode == "offline" {
	filesmanagement.CreateFile(cl.arguments[2])
	} else if cl.conMode == "online" {
		filesmanagement.ReqCreateFile(cl.arguments[2], myFile)
	}
	return cl.arguments[0], nil
}


func (cl *Schoise) DeleteFile() (string, error) {
	if cl.conMode == "offline" {
		filesmanagement.DeleteFile(cl.arguments[2])
	} else if cl.conMode == "online" {
		filesmanagement.ReqDeleteFile(cl.arguments[2], filesmanagement.Sfile{Name: cl.arguments[2]})
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) ReadFile() (string, error) {
	//fmt.Println("cl.conMode",cl.conMode)
	if cl.conMode == "offline" {
		filesmanagement.ReadFile(cl.arguments[2])
	} else if cl.conMode == "online" {
		//fmt.Println("cl.arguments[2]",cl.arguments[2])
		filesmanagement.ReqReadFile(cl.arguments[2])
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) RenameFile() (string, error) {
	if cl.conMode == "offline" {
	filesmanagement.RenameFile(cl.arguments[2], cl.arguments[3])
	} else if cl.conMode == "online" {
		filesmanagement.ReqRenameFile(cl.arguments[2], cl.arguments[3], filesmanagement.Sfile{Name: cl.arguments[2]})
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) ShowLogs() (string, error) {
	database.Condb()
	database.GetLogs()
	return cl.arguments[0], nil
}
