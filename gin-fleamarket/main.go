package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/middlewares"

	//"gin-fleamarket/models"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	router.Use(cors.Default())
	itemRouterWithAuth := router.Group("/items", middlewares.AuthMiddleware(authService))

	itemRouterWithAuth.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindByID)
	itemRouterWithAuth.POST("", itemController.Create)
	itemRouterWithAuth.PUT("/:id", itemController.Update)
	itemRouterWithAuth.DELETE("/:id", itemController.Delete)

	authRouter := router.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	return router
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	router := setupRouter(db)
	router.Run("localhost:8080")
}
