package main

import (
	"gin-mongo/configs"
	"gin-mongo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	 //run database
	configs.ConnectDB()

	//routes
    routes.UserRoute(router)

	router.Run("localhost:6000") 
}