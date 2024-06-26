package database

import (
	"go-auth-api/internal/application/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	properties := config.LoadEnv()
	dns := "host=" + *properties.DatabaseHost + " user=" + *properties.DatabaseUser + " password=" + *properties.DatabasePasswd + " dbname=" + *properties.DatabaseName + " port=" + *properties.DatabasePort

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Failed to connect on database")
	}

	migrationErr := db.AutoMigrate()

	if migrationErr != nil {
		panic("Failed to create tables")
	}

	return db
}
