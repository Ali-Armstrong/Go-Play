package main

import (
	"web/gin-jwt/configs"
	"web/gin-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//run database
	configs.ConnectDB()
	
	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")

}