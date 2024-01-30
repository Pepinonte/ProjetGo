package foldermanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func ReqDeleteFolder(name string, myFolder Sdossier) {	
	folder := deleteFolder(name, myFolder)
	fmt.Println("dossier supprimé",folder)
}

func deleteFolder(name string, monDossier Sdossier) Sdossier{
	jsonReq, _ := json.Marshal(monDossier)
	url := server + "/Folder/" + name
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonReq))
	if err != nil {
			log.Fatalln(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			log.Fatalln(err)
	}
	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 		log.Fatalln(resp.StatusCode)
	// }

	var rep Sdossier
	json.NewDecoder(resp.Body).Decode(&rep)

	var children []Children
	for _, child := range rep.Children {
			children = append(children, Children{Name: child.Name})
	}
	return Sdossier{Name: rep.Name, Children: children}
}