package config

import (
	"fmt"
	"github.com/moverq1337/JWTAuth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error
	LoadEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("pghost"), os.Getenv("pguser"), os.Getenv("pgpass"), os.Getenv("pgdb"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("not connected to database:", err)
		panic(err)
	}
	log.Println("connected to database")
	DB.AutoMigrate(models.User{})
}
