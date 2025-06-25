package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env")
		return
	}
	log.Print(".env loaded")
}
