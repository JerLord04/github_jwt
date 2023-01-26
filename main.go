package main

import (
	"example.com/m/controllers"
	initializers "example.com/m/initializer"
	"example.com/m/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.Run()
}
