package main

import (
	"fmt"
	"os"

	"github.com/zetsux/gin-gorm-template-clean/api/v1/controller"
	"github.com/zetsux/gin-gorm-template-clean/api/v1/router"
	"github.com/zetsux/gin-gorm-template-clean/common/middleware"
	"github.com/zetsux/gin-gorm-template-clean/config"
	"github.com/zetsux/gin-gorm-template-clean/core/repository"
	"github.com/zetsux/gin-gorm-template-clean/core/service"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		db   = config.DBSetup()
		jwtS = service.NewJWTService()

		userR = repository.NewUserRepository(db)
		userS = service.NewUserService(userR)
		userC = controller.NewUserController(userS, jwtS)

		fileC = controller.NewFileController()
	)

	defer config.DBClose(db)

	// Setting Up Server
	server := gin.Default()
	server.Use(
		middleware.CORSMiddleware(),
	)

	// Setting Up Routes
	router.UserRouter(server, userC, jwtS)
	router.FileRouter(server, fileC)

	// Running in localhost:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := server.Run(":" + port)
	if err != nil {
		fmt.Println("Server failed to start: ", err)
		return
	}
}
