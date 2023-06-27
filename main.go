package main

import (
	routes "go-weather/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()

	router := gin.Default()
	ConnectDB()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	router.POST("/signup", routes.Signup)
	router.POST("/login", routes.Login)
	router.POST("/profile", routes.Profile)

	router.Run()
}
