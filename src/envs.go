package src

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() (string, string) {
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Error loading .env file:", err)
	}

	username := os.Getenv("TWITCH_USERNAME")
	if username == "" {
		log.Fatal("Missing environment variable TWITCH_USERNAME")
	}

	oauthToken := os.Getenv("TWITCH_OAUTH_TOKEN")
	if oauthToken == "" {
		log.Fatal("Missing environment variable TWITCH_OAUTH_TOKEN")
	}

	return username, oauthToken
}
