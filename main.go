package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	LoadEnv()

	router := gin.Default()
	instance := ConnectDB()

	router.Use(useDatabase(instance.DB))

	fmt.Print(instance.DB.Stats().WaitDuration.String())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server is up and running",
		})
	})

	router.Run()
}

func useDatabase(instance *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		c.Set("db", instance)

		// before request

		c.Next()
	}
}
