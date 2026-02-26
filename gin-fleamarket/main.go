package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "T-shirt", Price: 1000, Description: "A nice T-shirt", SoldOut: false},
		{ID: 2, Name: "Jeans", Price: 2000, Description: "A pair of jeans", SoldOut: true},
		{ID: 3, Name: "Sneakers", Price: 3000, Description: "A pair of sneakers", SoldOut: true},
	}

	itemReposiory := repositories.NewItemMemoryRepository(items)
	itemService := services.NewItemService(itemReposiory)
	itemController := controllers.NewItemController(itemService)

	router := gin.Default()
	router.GET("/items", itemController.FindAll)
	router.GET("/items/:id", itemController.FindByID)
	router.POST("/items", itemController.Create)
	router.PUT("/items/:id", itemController.Update)
	router.DELETE("/items/:id", itemController.Delete)
	router.Run("localhost:8080")
}
