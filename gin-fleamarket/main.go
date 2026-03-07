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
	authRouter.POST("/login", authController.Login)

	router.Run("localhost:8080")
}
