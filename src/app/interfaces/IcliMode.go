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
	CreateFile(path string) (string, error)
	DeleteFile(path string) (string, error)
	ReadFile(path string) (string, error)
	RenameFile(path string) (string, error)
	ShowLogs() (string, error)
}

// TODO: try no define the struct here and use the one from app/structures/Schoise.go
type Schoise struct {
	mode      string
	arguments []string
	conMode	 string
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
		cl.CreateFile(st.arguments[2])
	case "deleteFile":
		cl.DeleteFile(st.arguments[2])
	case "readFile":
		cl.ReadFile(st.arguments[2])
	case "renameFile":
		cl.RenameFile(st.arguments[2])
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

func (cl *Schoise) CreateFile(path string) (string, error) {
	filesmanagement.CreateFile(cl.arguments[2])
	return cl.arguments[0], nil
}

func (cl *Schoise) DeleteFile(path string) (string, error) {
	filesmanagement.DeleteFile(cl.arguments[2])
	return cl.arguments[0], nil
}

func (cl *Schoise) ReadFile(path string) (string, error) {
	filesmanagement.ReadFile(cl.arguments[2])
	return cl.arguments[0], nil
}

func (cl *Schoise) RenameFile(path string) (string, error) {
	filesmanagement.RenameFile(cl.arguments[2], cl.arguments[3])
	return cl.arguments[0], nil
}

func (cl *Schoise) ShowLogs() (string, error) {
	database.Condb()
	database.GetLogs()
	return cl.arguments[0], nil
}
