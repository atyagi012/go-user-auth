package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort      string
	Postgres_URL string
	SecretKey    string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		log.Panic("Error loading .env file")
	}

	e.AppPort = loadString("APP_PORT", "8283")
	fmt.Println("APP_PORT:", e.AppPort)

	e.Postgres_URL = loadString("POSTGRES_URL", "postgres://postgres:adminPassword@localhost:5432/GO_USER_JWT")
	fmt.Println("POSTGRES_URL:", e.Postgres_URL)

	e.SecretKey = loadString("SECERET", "")
	fmt.Println("SECRET_KEY:", e.SecretKey)
}

var Config envConfig

func init() {
	fmt.Println("Initializing config package...")
	Config.LoadConfig()
}

func loadString(key string, fallback string) string {
	envValue, ok := os.LookupEnv(key)
	if !ok {
		log.Panic("APP_PORT not set in .env file")
		return fallback
	}
	return envValue
}
