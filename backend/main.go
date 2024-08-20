package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var items []Item

func main() {
	items = []Item{
		{ID: 1, Name: "Galactic Goggles"},
		{ID: 2, Name: "Meteor Muffins"},
		{ID: 3, Name: "Alien Antenna Kit"},
		{ID: 4, Name: "Starlight Lantern"},
		{ID: 5, Name: "Quantum Quill"},
	}

	router := gin.Default()
	router.GET("/", greet)
	router.GET("/item/:id", getSingleItem)
	router.GET("/items", getItems)
	router.POST("/items", addItem)
	router.HEAD("/healthcheck", healthcheck)

	router.Run()
}

func greet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome, Go navigator, to the Anythink cosmic catalog.")
}

type Item struct {
    ID   int    `json:"id" uri:"id" `
    Name string `json:"name"`
	ViewCount int `json:"-"`
}

func findItemById(items *[]Item, id int) *Item {
	for i := range *items {
		if (*items)[i].ID == id {
			return &(*items)[i]
		}
	}
	return nil
}

func getSingleItem(c *gin.Context) {

	var input Item

	if err := c.BindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	singleItem := findItemById(&items, input.ID)

	if singleItem == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No item found"})
		return
	}

	go func() {
		singleItem.ViewCount = singleItem.ViewCount + 1
	}()

	c.IndentedJSON(http.StatusOK, singleItem)
}

func addItem(c *gin.Context) {
	lastItem := items[len(items)-1]

	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem.ID = lastItem.ID + 1

	items = append(items, newItem)

	c.IndentedJSON(http.StatusOK, newItem)
}

func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, items)
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
