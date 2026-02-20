package main

import (
	"gin-fleamarket/models"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "T-shirt", Prince: 1000, Description: "A nice T-shirt", SoldOut: false},
		{ID: 2, Name: "Jeans", Prince: 2000, Description: "A pair of jeans", SoldOut: true},
		{ID: 3, Name: "Sneakers", Prince: 3000, Description: "A pair of sneakers", SoldOut: true},
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("localhost:8080")
}
