package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", greet)
	router.GET("/items", items)
	router.HEAD("/healthcheck", healthcheck)

	router.Run()
}

func greet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome, Go navigator, to the Anythink cosmic catalog.")
}

func items(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, []map[string]interface{}{
		{
			"id": 1,
			"name":"Galactic Goggles",
		},
		{
			"id": 2,
			"name":"Meteor Muffins",
		},
		{
			"id": 3,
			"name":"Alien Antenna Kit",
		},
		{
			"id": 4,
			"name":"Starlight Lantern",
		},
		{
			"id": 5,
			"name":"Quantum Quill",
		},
	})
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
