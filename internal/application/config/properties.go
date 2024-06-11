package config

import (
	"github.com/joho/godotenv"
	"go-auth-api/internal/domain/types"
	"os"
)

func LoadEnv() *types.EnvProperties {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &types.EnvProperties{
		DatabaseHost:   os.Getenv("DATABASE_HOST"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		DatabasePasswd: os.Getenv("DATABASE_PASSWD"),
		DatabaseUser:   os.Getenv("DATABASE_USER"),
		DatabasePort:   os.Getenv("DATABASE_PORT"),
	}

}
