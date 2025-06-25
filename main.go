package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/JWTAuth/config"
	"github.com/moverq1337/JWTAuth/controllers"
)

func main() {
	config.Connect()
	r := gin.Default()
	fmt.Println("hello")
	r.POST("/registration", controllers.UserRegistation)
	r.POST("/login", controllers.UserLogin)
	r.Run()
}
