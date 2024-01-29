package foldermanagement

import (
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

func ReqCreateFolder(name string) {	
	createFolder(name)
}

func createFolder(name string) {
	url := server + "/createFolder/" + name
	resp, err := http.Post(url, "application/json; charset=utf-8", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalln(resp.StatusCode)
	}
}