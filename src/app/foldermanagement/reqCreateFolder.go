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
	// albums := getAll()
	// log.Println(albums)
	// album := getOne("1")
	// log.Println(album)
	createFolder(name)
}

// func getAll() []Album {
// 	url := server
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalln(resp.StatusCode)
// 	}

// 	var albums []Album
// 	bodybytes, _ := io.ReadAll(resp.Body)
// 	json.Unmarshal(bodybytes, &albums)
// 	return albums
// }

// func getOne(id string) Album {
// 	url := server + "/" + id
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalln(resp.StatusCode)
// 	}

// 	var album Album
// 	bodybytes, _ := io.ReadAll(resp.Body)
// 	json.Unmarshal(bodybytes, &album)
// 	return album
// }

// func post(album Album) Album {
// 	jsonReq, err := json.Marshal(album)
// 	resp, err := http.Post(server, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusCreated {
// 		log.Fatalln(resp.StatusCode)
// 	}

// 	var posted Album
// 	bodybytes, _ := io.ReadAll(resp.Body)
// 	json.Unmarshal(bodybytes, &posted)
// 	return posted
// }

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