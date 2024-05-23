package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                 string
	Port                 string
	DbUser               string
	DbPassword           string
	DbAddress            string
	DbName               string
	JwtSecret            string
	JwtExpirationInHours int64
}

var Envs = loadConfig()

func loadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading enviroment variables from file. err: %s", err.Error())
	}

	dbHost := getEnvStr("DB_HOST", "localhost")
	dbPort := getEnvStr("DB_PORT", "3306")

	return Config{
		Host:                 getEnvStr("HOST", "http://localhost"),
		Port:                 getEnvStr("PORT", ""),
		DbUser:               getEnvStr("DB_USER", ""),
		DbPassword:           getEnvStr("DB_PASSWORD", ""),
		DbAddress:            fmt.Sprintf("%s:%s", dbHost, dbPort),
		DbName:               getEnvStr("DB_NAME", ""),
		JwtSecret:            getEnvStr("JWT_SECRET", ""),
		JwtExpirationInHours: getEnvInt("JWT_EXPIRATION_IN_HOURS", 24), //24 hours
	}
}

func getEnvStr(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok && fallback == "" {
		log.Fatalf("env variable %s not found", key)
	}

	if !ok && fallback != "" {
		return fallback
	}

	return string(value)
}

func getEnvInt(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)

	if ok {
		intValue, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			log.Fatalf("error parsing env variable %s to int", key)
		}

		return intValue
	}

	return fallback
}
