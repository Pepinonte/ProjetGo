package foldermanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



func ReqRenameFolder(lname string, nname string,  myFolder Sdossier){	
	dossier := renameFolder(lname, nname, myFolder)
	fmt.Println("dossier modifié",dossier)
}

func renameFolder(ln string, nn string, monDossier Sdossier) Sdossier{
	jsonReq, _ := json.Marshal(monDossier)
	url := server + "/Folder/" + ln + "/" + nn
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			log.Fatalln(err)
	}
	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusCreated {
	// 	log.Fatalln(resp.StatusCode)
	// }
	if resp.StatusCode != http.StatusOK {
		log.Fatalln(resp.Body)
	}
	var rep Sdossier
	json.NewDecoder(resp.Body).Decode(&rep)

	var children []Children
	for _, child := range rep.Children {
			children = append(children, Children{Name: child.Name})
	}
	return Sdossier{Name: rep.Name, Children: children}
}