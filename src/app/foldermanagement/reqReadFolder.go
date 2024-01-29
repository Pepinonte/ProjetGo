package foldermanagement

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Sdossier struct {
	ID    		string `json:"id"`
	name		string `json:"name"`
	children	[]os.DirEntry `json:"children"`
}

func ReqReadFolder(name string) {	
	readFolder(name)
}

func readFolder(name string) {
	url := server + "/readFolder/" + name
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalln(resp.Body)
	}
	var rep []Sdossier
	bodybytes, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bodybytes, &rep)
	fmt.Println("dsdsddddddddd",rep)

}