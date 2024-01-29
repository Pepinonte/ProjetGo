package api

import (
	"app/database"
	"app/filesmanagement"
	"app/foldermanagement"

	"github.com/gin-gonic/gin"
)

func RunListener() {
	router := gin.Default()
	router.GET("showLogs", showLogs)
	router.POST("/createFolder/:name", createFolder)
	router.DELETE("/deleteFolder/:name", deleteFolder)
	router.GET("/readFolder/:name", readFolder)
	router.PUT("/renameFolder/:lname/:nname", renameFolder)
	router.POST("/createFile/:name", createFile)
	router.DELETE("/deleteFile/:name", deleteFile)
	router.PUT("/modifyFile/:name/:content", modifyFile)
	router.GET("/readFile/:name", readFile)
	router.PUT("/renameFile/:lname/:nname", renameFile)
	router.Run("localhost:8080")
}

func createFolder(c *gin.Context) {
	name := c.Param("name")
	foldermanagement.CreateFolder(name)
}

func deleteFolder(c *gin.Context) {
	name := c.Param("name")
	foldermanagement.DeleteFolder(name)
}

func readFolder(c *gin.Context) {
	name := c.Param("name")
	foldermanagement.ReadFolder(name)
}

func renameFolder(c *gin.Context) {
	lname := c.Param("lname")
	nname := c.Param("nname")
	foldermanagement.RenameFolder(lname, nname)
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
