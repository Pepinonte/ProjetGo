package filesmanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func ReqDeleteFile(name string, myFile Sfile) {	
	fmt.Println("name",name)
	file := deleteFile(name, myFile)
	log.Println("Fichier supprimé :",file)
}

func deleteFile(name string, myFile Sfile) Sfile{
	jsonReq, _ := json.Marshal(myFile)
	url := server + "/deleteFile/" + name
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

	if resp.StatusCode != http.StatusOK {
			log.Fatalln(resp.StatusCode)
	}

	var result Sfile
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}