package controllers

import "github.com/gin-gonic/gin"

func ProtectedHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "protected",
	})
}
