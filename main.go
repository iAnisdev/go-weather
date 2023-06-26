package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()

	router := gin.Default()

	DB := ConnectDB()

	fmt.Print(DB.DB.Stats().WaitDuration.String())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	router.Run()
}
