package main

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	"net/http"
)

type testModel struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main() {
	// creating a gin router
	r := gin.Default()

	r.GET("/test/:name", func(c *gin.Context) {
		name := c.Param("name")
		// c.String(http.StatusOK, "Hello %s", name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	r.Run(":8080")
}