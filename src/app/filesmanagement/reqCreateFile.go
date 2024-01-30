package filesmanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



func ReqCreateFile(name string, myFile Sfile) {	
	file := createFile(name, myFile)
	fmt.Println("Fichier Crée :",file)
}

func createFile(name string, myFile Sfile) Sfile {
	jsonReq, _ := json.Marshal(myFile)
	
	
    url := server + "/createFile/" + name
	
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	

	if resp.StatusCode != http.StatusCreated {
		log.Fatalln(resp.StatusCode)
	}
	var result Sfile
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	myFile.Name = name
	myFile.Content = result.Content

	return myFile

}