package main

import (
	"msusers/controller"
	"msusers/database"
	"msusers/middleware"
	"msusers/models"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.Bank{})
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Wallet{})
}

func loadEnv() {
	err := godotenv.Load(".Env")
	if err != nil  {
		log.Fatal("error loading Env File")
	}
}

func serverApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())

	protectedRoutes.POST("/users", controller.CreateUser)
	protectedRoutes.GET("/users", controller.ListAllUser)
	protectedRoutes.GET("/users/:id",controller.GetUserDetails)
	protectedRoutes.PUT("/users/:id", controller.UpdatedUser)
	protectedRoutes.DELETE("users/:id", controller.DeleteUser)

	protectedRoutes.POST("/banks", controller.SaveBank)
	protectedRoutes.GET("/bank", controller.ListBank)
	protectedRoutes.GET("/banks", controller.ListAllBank)
	protectedRoutes.GET("/banks/:id",controller.GetBankDetailsById)
	protectedRoutes.PUT("/banks/:id", controller.UpdatedBank)
	protectedRoutes.DELETE("/banks/:id", controller.DeleteBank)

	protectedRoutes.POST("/wallets", controller.SaveWallet)
	protectedRoutes.GET("/wallet", controller.ListWallet)
	protectedRoutes.GET("/wallets", controller.ListAllWallets)
	protectedRoutes.GET("/wallets/:id", controller.GetWalletsById)
	protectedRoutes.PUT("/wallets/:id", controller.UpdatedWallet)
	protectedRoutes.DELETE("/wallets/:id", controller.DeleteWallet)

	router.Run(":8000")
	fmt.Println("Server running at port 8000")
}