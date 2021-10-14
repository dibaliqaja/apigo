package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"dibaliqaja/apigo/config"
	"dibaliqaja/apigo/controller"
)

var (
	db 			   *gorm.DB 				 = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main () {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	r.Run()
}