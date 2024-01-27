package api

import (
	"app/foldermanagement"

	"github.com/gin-gonic/gin"
)

func RunListener() {
	router := gin.Default()
	router.GET("/createFolder/:name", createFolder)
	router.Run("localhost:8080")
}


func createFolder(c *gin.Context) {
	name := c.Param("name")
	foldermanagement.CreateFolder(name)
		
}
