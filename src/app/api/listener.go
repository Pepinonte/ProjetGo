package api

import (
	"app/database"
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

func showLogs(c *gin.Context) {
	database.GetLogs()
}
