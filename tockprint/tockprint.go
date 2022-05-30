package tockprint

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func TockInit() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func P() string {
	n := os.Getenv("name")
	return n
}