package config

import (
	"context"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/utils"
	"os"
)

func LoadEnv() *types.EnvProperties {
	err := godotenv.Load(".env")
	if err != nil {
		return nil
	}
	return &types.EnvProperties{
		DatabaseHost:   utils.ToPointer(os.Getenv("DATABASE_HOST")),
		DatabaseName:   utils.ToPointer(os.Getenv("DATABASE_NAME")),
		DatabasePasswd: utils.ToPointer(os.Getenv("DATABASE_PASSWD")),
		DatabaseUser:   utils.ToPointer(os.Getenv("DATABASE_USER")),
		DatabasePort:   utils.ToPointer(os.Getenv("DATABASE_PORT")),
	}

}

func GetRequestId(ctx context.Context) string {
	reqId, _ := ctx.Value("request_id").(string)

	return reqId
}

func SetRequestId(ctx context.Context) context.Context {

	id, _ := uuid.NewUUID()
	ctx = context.WithValue(ctx, "request_id", id.String())

	return ctx
}
