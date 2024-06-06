package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the GET request at "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Define a route for the GET request at "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define a route for the POST request at "/post"
	router.POST("/post", func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
			Age  int    `json:"age" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name": json.Name,
			"age":  json.Age,
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
