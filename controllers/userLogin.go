package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/JWTAuth/config"
	"github.com/moverq1337/JWTAuth/models"
)

func UserLogin(c *gin.Context) {
	var LoginUser models.User
	if err := c.ShouldBindJSON(&LoginUser); err != nil {
		c.JSON(400, gin.H{
			"message": "not json",
		})
		return
	}
	if LoginUser.Name == "" || LoginUser.Password == "" {
		c.JSON(400, gin.H{
			"message": "complete fields",
		})
	}
	var dbUser models.User //создаем нового для хранения данных из бд котороое отдастся указателем в эту переменную
	if err := config.DB.Where("name = ?", LoginUser.Name).First(&dbUser).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "incorrect username or password",
		})
		fmt.Println(err)
		return
	}
	if !CheckPasswordHash(LoginUser.Password, dbUser.Password) {
		c.JSON(400, gin.H{
			"message": "incorrect username or password",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "got it",
	})
}
