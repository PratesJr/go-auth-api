package config

import (
	"github.com/joho/godotenv"
	"go-auth-api/internal/domain/interfaces"
	"os"
)

func LoadEnv() *interfaces.EnvProperties {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &interfaces.EnvProperties{
		DatabaseHost:   os.Getenv("DATABASE_HOST"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		DatabasePasswd: os.Getenv("DATABASE_PASSWD"),
		DatabaseUser:   os.Getenv("DATABASE_USER"),
		DatabasePort:   os.Getenv("DATABASE_PORT"),
	}

}
