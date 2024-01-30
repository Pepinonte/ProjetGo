package interfaces

import (
	"app/database"
	"app/filesmanagement"
	"app/foldermanagement"
	"errors"
	"fmt"
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

type Children struct {
	Name string
}

type Sdossier struct {
	ID    		string `json:"id"`
	Name		string `json:"name"`
	Children	[]Children `json:"children"`
}

func IcliModeMain(conMode string,arguments []string) {
	mycli := Schoise{arguments[1], arguments, conMode}
	execute(&mycli ,mycli)
}

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
	if len(cl.arguments) < 3 {
		fmt.Println("path vide")
		return "", errors.New("path vide")
	}
	myFolder, err, _ := foldermanagement.CreateFolder(cl.arguments[2])
	if cl.conMode == "offline" {
		
		// dos,err,_ := foldermanagement.CreateFolder(cl.arguments[2])
		if err != nil {
			fmt.Println("une erreur c'est produite",err)
			return "", err
		}
		fmt.Println("dossier crée",myFolder)
	} else if cl.conMode == "online" {
		foldermanagement.ReqCreateFolder(cl.arguments[2], myFolder)
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) DeleteFolder() (string, error) {
	if len(cl.arguments) < 3 {
		fmt.Println("path vide")
		return "", errors.New("path vide")
	}

	myFolder, _, _ := foldermanagement.DeleteFolder(cl.arguments[2])
	if cl.conMode == "offline" {
		// foldermanagement.DeleteFolder(cl.arguments[2])
	} else if cl.conMode == "online" {
		foldermanagement.ReqDeleteFolder(cl.arguments[2], myFolder)
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) ReadFolder() (string, error) {

	if len(cl.arguments) < 3 {
		fmt.Println("path vide")
		return "", errors.New("path vide")
	}

	myFolder, _, _ := foldermanagement.ReadFolder(cl.arguments[2])
	if cl.conMode == "offline" {
		// foldermanagement.ReadFolder(cl.arguments[2])
	} else if cl.conMode == "online" {
		foldermanagement.ReqReadFolder(cl.arguments[2], myFolder)
	}
	return cl.arguments[0], nil
}

func (cl *Schoise) RenameFolder() (string, error) {
	if len(cl.arguments) < 3 {
		fmt.Println("path vide")
		return "", errors.New("il manque des arguments")
	} 

	myFolder, _, _ := foldermanagement.RenameFolder(cl.arguments[2], cl.arguments[3])
	if cl.conMode == "offline" {
		// foldermanagement.RenameFolder(cl.arguments[2], cl.arguments[3])
	} else if cl.conMode == "online" {
		foldermanagement.ReqRenameFolder(cl.arguments[2], cl.arguments[3], myFolder)
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
