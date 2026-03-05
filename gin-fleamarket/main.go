package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"

	//"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	//items := []models.Item{
	//{ID: 1, Name: "T-shirt", Price: 1000, Description: "A nice T-shirt", SoldOut: false},
	//{ID: 2, Name: "Jeans", Price: 2000, Description: "A pair of jeans", SoldOut: true},
	//{ID: 3, Name: "Sneakers", Price: 3000, Description: "A pair of sneakers", SoldOut: true},
	//}

	//itemReposiory := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	itemRouter := router.Group("/items")
	itemRouter.GET("/items", itemController.FindAll)
	itemRouter.GET("/items/:id", itemController.FindByID)
	itemRouter.POST("/items", itemController.Create)
	itemRouter.PUT("/items/:id", itemController.Update)
	itemRouter.DELETE("/items/:id", itemController.Delete)

	authRouter := router.Group("/auth")
	authRouter.POST("/signup", authController.Signup)

	router.Run("localhost:8080")
}
