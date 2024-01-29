package foldermanagement

import (
	"log"
	"net/http"
)


func ReqDeleteFolder(name string) {	
	deleteFolder(name)
}

func deleteFolder(name string) {
	url := server + "/deleteFolder/" + name
	req, err := http.NewRequest("DELETE", url, nil)
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
}