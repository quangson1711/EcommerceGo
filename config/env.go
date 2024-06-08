package config

import (
	"github.com/lpernett/godotenv"
	"os"
	"strconv"
)

type Config struct {
	PublicHost            string
	Port                  string
	DBUser                string
	DBPassword            string
	DBAddress             string
	DBName                string
	JWTExpirationInSecond int
	JWTSecret             string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:            os.Getenv("PUBLIC_HOST"),
		Port:                  os.Getenv("PORT"),
		DBUser:                os.Getenv("DB_USER"),
		DBPassword:            os.Getenv("DB_PASSWORD"),
		DBAddress:             os.Getenv("DB_ADDRESS"),
		DBName:                os.Getenv("DB_NAME"),
		JWTExpirationInSecond: getEnvAsInt("JWT_EXPIRATION_IN_SECOND", 60*60*24*7),
		JWTSecret:             os.Getenv("JWT_SECRET"),
	}
}

func getEnvAsInt(key string, fallback int) int {
	v, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return v
}
