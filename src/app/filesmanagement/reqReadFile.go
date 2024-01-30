package filesmanagement

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)


type Sfile struct {
	ID    		string `json:"id"`
	Name		string `json:"name"`
	Content	string `json:"content"`
}

const server = "http://localhost:3307"


func ReqReadFile(name string) {	
	test := readFile(name)
	fmt.Println("Fichier lu :",test)
}

func readFile(name string) Sfile {
    url := server + "/readFile/" + name
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Fatalln(bodyString)
	}
    var result string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
 		log.Fatalln(err)
	}

	var myFile Sfile
	myFile.Name = name
	myFile.Content = result


	return myFile

}