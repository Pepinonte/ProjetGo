package foldermanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Schoise struct {
	ID    		string `json:"id"`
	date			string `json:"date"`
	mode      string `json:"mode"`
	module 		string	`json:"module"`
	function 	string	`json:"function"`
	arguments []string	`json:"arguments"`
	conMode	  string	`json:"conMode"`
	output	  string	`json:"output"`
}

const server = "http://localhost:8080"

func ReqCreateFolder(name string, myFolder Sdossier){	
	test := createFolder(name, myFolder)
	fmt.Println("dossier crée",test)
}

func createFolder(name string, monDossier Sdossier) Sdossier{
	jsonReq, _ := json.Marshal(monDossier)
	url := server + "/createFolder/" + name
	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalln(resp.StatusCode)
	}
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