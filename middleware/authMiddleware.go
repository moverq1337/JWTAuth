package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/JWTAuth/utils"
	"log"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	res := c.GetHeader("Authorization")
	tokenString := strings.Split(res, " ")
	if len(tokenString) != 2 || tokenString[0] != "Bearer" {
		c.JSON(400, gin.H{
			"error": "token is incorrect type",
		})
		c.Abort()
		return
	}

	token := tokenString[1]
	if token == "" {
		c.JSON(400, gin.H{
			"error": "token is null",
		})
		c.Abort()
		return
	}
	if err := utils.VerifyToken(token); err != nil {
		c.JSON(400, gin.H{
			"error": "token is not valid",
		})
		log.Println(err)
		c.Abort()
		return
	}
	c.Next()
}
