package foldermanagement

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Children struct{
	Name string
}

type Sdossier struct {
	ID    		string `json:"id"`
	Name		string `json:"name"`
	Children	[]Children `json:"children"`
}



func ReqReadFolder(name string, myFolder Sdossier) {	
	test := readFolder(name)
	fmt.Println("dossier affiché",test)
}

func readFolder(name string) Sdossier {
    url := server + "/readFolder/" + name
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

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