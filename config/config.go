package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("Error loading .env file\n")
	}

	v := os.Getenv(key)
	if v == "" {
		fmt.Printf("Env variable \"%s\" not defined.\n", key)
	}

	return v
}
