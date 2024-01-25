package main

import (
	"app/foldermanagement"
	"fmt"
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
		createFile(os.Args[2], os.Args[3])
	case "deleteFile":
		if len(os.Args) >= 4 {
			deleteFile(os.Args[2], os.Args[3])
		} else {
			deleteFile(os.Args[2], "")
		}	
	case "readFile":
		if len(os.Args) >= 4 {
			readFile(os.Args[2], os.Args[3])
		} else {
			readFile("", os.Args[2])
		}
	case "modifyFile":
		modifyFile(os.Args[2], os.Args[3])
	case "renameFile":
		renameFile(os.Args[2], os.Args[3])
	}

	


}


// folder functions



// end folder functions

// file functions
func createFile(path string, name string) {
	_, err := os.Create(path + "/" + name)
	if err != nil {
		fmt.Println("impossible de creer le fichier")
		return
	}
}

func deleteFile(path string, name string) {
	if path == "" {
		path = "."
	}
	err := os.Remove(path + "/" + name)
	if err != nil {
		fmt.Println("Impossible de supprimer le fichier")
		return
	}
}


func readFile(path string, name string) {
	if path == "" {
		path = "."
	}

	file, error := os.Open(path + "/" + name)
	if error != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return
	}

	defer file.Close()

	stat, error := file.Stat()
	if error != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return
	}

	bs := make([]byte, stat.Size())
	_, error = file.Read(bs)
	if error != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func modifyFile(name string, content string) {
	file, error := os.OpenFile(name, os.O_RDWR|os.O_TRUNC, 0644)
	if error != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		return
	}

	defer file.Close()

	file.WriteString(content)
}

func renameFile(name string, newName string) {
    err := os.Rename(name, newName)
    if err != nil {
        fmt.Println("Impossible de renommer le fichier")
        return
    }
}

