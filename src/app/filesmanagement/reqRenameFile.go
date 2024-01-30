package filesmanagement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



func ReqRenameFile(name string,nname string, myFile Sfile) {	
	file := renameFile(name,nname, myFile)
	fmt.Println("Fichier renommé :",file)
}

func renameFile(name string,nname string, myFile Sfile) Sfile {
	jsonReq, _ := json.Marshal(myFile)
	
	
    url := server + "/renameFile/" + name + "/" + nname
	
    resp, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	

	
	var result Sfile
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	myFile.Name = nname

	return myFile

}