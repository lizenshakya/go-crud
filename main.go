package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lizenshakya/go-crud/controllers"
	"github.com/lizenshakya/go-crud/initializers"
	"github.com/lizenshakya/go-crud/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Login)
	r.GET("/verify", middleware.RequireAuth, controllers.VerifyToken)

	r.Run() // listen and serve on 0.0.0.0:8080
}
