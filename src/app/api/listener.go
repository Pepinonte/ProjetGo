package api

import (
	"app/database"
	"app/filesmanagement"
	"app/foldermanagement"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Children struct {
	Name string
}

type Sdossier struct {
	ID    		string `json:"id"`
	Name		string `json:"name"`
	Children	[]Children `json:"children"`
}


func RunListener() {
	//TODO: penser a adapter le nom des routes pour s'adapter au REST
	router := gin.Default()
	router.GET("showLogs", showLogs)
	router.POST("/Folder/:name", createFolder)
	router.DELETE("/Folder/:name", deleteFolder)
	router.GET("/Folder/:name", readFolder)
	router.PUT("/Folder/:lname/:nname", renameFolder)
	router.POST("/createFile/:name", createFile)
	router.DELETE("/deleteFile/:name", deleteFile)
	router.PUT("/modifyFile/:name/:content", modifyFile)
	router.GET("/readFile/:name", readFile)
	router.PUT("/renameFile/:lname/:nname", renameFile)
	router.Run("localhost:8080")
}

func createFolder(c *gin.Context) {
	name := c.Param("name")
	_, err, data := foldermanagement.CreateFolder(name)

	// if len(data) < 3 {
	// 	fmt.Println("path vide")
	// 	return
	// }

	if err != nil {
		fmt.Println("err", err)
	}
	var children []Children
	for _, childName := range data {
			children = append(children, Children{Name: childName})
	}

	var myFolder = Sdossier{Name: name, Children: children}
	// fmt.Println("myFolder",myFolder)
	c.IndentedJSON(http.StatusOK, myFolder)
}

func deleteFolder(c *gin.Context) {
	name := c.Param("name")
	_, _, data := foldermanagement.DeleteFolder(name)

	var children []Children
    for _, childName := range data {
        children = append(children, Children{Name: childName})
    }

    var myFolder = Sdossier{Name: name, Children: children}
    // fmt.Println("myFolder",myFolder)
    c.IndentedJSON(http.StatusOK, myFolder)
}

func readFolder(c *gin.Context) {
    name := c.Param("name")
    _,_,data := foldermanagement.ReadFolder(name)

    var children []Children
    for _, childName := range data {
        children = append(children, Children{Name: childName})
    }

    var myFolder = Sdossier{Name: name, Children: children}
    // fmt.Println("myFolder",myFolder)
    c.IndentedJSON(http.StatusOK, myFolder)
}

func renameFolder(c *gin.Context) {
	lname := c.Param("lname")
	nname := c.Param("nname")
	_,_, data := foldermanagement.RenameFolder(lname, nname)
	// fmt.Println("data",data)

	var children []Children
    for _, childName := range data {
        children = append(children, Children{Name: childName})
    }

    var myFolder = Sdossier{Name: nname, Children: children}
    // fmt.Println("myFolder",myFolder)
    c.IndentedJSON(http.StatusOK, myFolder)
}

func createFile(c *gin.Context) {
	name := c.Param("name")
	filesmanagement.CreateFile(name)
}

func deleteFile(c *gin.Context) {
	name := c.Param("name")
	//fmt.Println("je suis ici")
	filesmanagement.DeleteFile(name)
}

func modifyFile(c *gin.Context) {
	name := c.Param("name")
	content := c.Param("content")
	filesmanagement.ModifyFile(name, content)
}

func readFile(c *gin.Context) {
	name := c.Param("name")
	filesmanagement.ReadFile(name)
}

func renameFile(c *gin.Context) {
	lname := c.Param("lname")
	nname := c.Param("nname")
	filesmanagement.RenameFile(lname, nname)
}

func showLogs(c *gin.Context) {
	database.GetLogs()
}
