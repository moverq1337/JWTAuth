package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/JWTAuth/utils"
	"log"
)

type TokenRequest struct {
	Token string `json:"token"`
}

func AuthMiddleware(c *gin.Context) {
	var req TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "token is not defined",
		})
		log.Println(err)
		c.Abort()

	}
	if req.Token == "" {
		c.JSON(400, gin.H{
			"error": "token is null",
		})
		c.Abort()

	}
	if err := utils.VerifyToken(req.Token); err != nil {
		c.JSON(400, gin.H{
			"error": "token is not valid",
		})
		log.Println(err)
		c.Abort()
	}
	c.Next()
}
