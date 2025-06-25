package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moverq1337/JWTAuth/config"
	"github.com/moverq1337/JWTAuth/models"
)

func UserRegistation(c *gin.Context) {
	var NewUser models.User
	if err := c.ShouldBindJSON(&NewUser); err != nil {
		fmt.Println("Не json")
		return
	}
	if NewUser.Name == "" || NewUser.Password == "" {
		fmt.Println("Заполните данные")
		c.JSON(400, gin.H{
			"message": "Заполните данные",
		})
		return
	}
	hash, err := HashPassword(NewUser.Password)
	if err != nil {
		fmt.Println("Не удалось захэшировать пароль")
		return
	}
	if CheckPasswordHash(NewUser.Password, hash) != true {
		fmt.Println("Пароль не соответсвует хэшу")
		return
	}
	if err = config.DB.Create(&models.User{Name: NewUser.Name, Password: hash}).Error; err != nil {
		fmt.Println("Не удалось создать человека")
		return
	}
	c.JSON(200, gin.H{
		"message": "done",
	})

}
