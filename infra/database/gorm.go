package database

import (
	"code-space-backend-api/env"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

func Load() {
	Connect()

	MigrateModels()
	SeedTables()

	currentPath, _ := os.Getwd()
	RunSQLScriptsInFolder(fmt.Sprintf("%s/%s", currentPath, "infra/database/triggers"))
}

func Connect() *gorm.DB {
	if instance != nil {
		return instance
	}

	newInstance, err := Open(env.Database.DSN)
	if err != nil {
		log.Fatal("error on get DB instance ", err)
	}

	instance = newInstance
	return instance
}

func Open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Close() {
	err := CloseDB(instance)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB(gormDB *gorm.DB) error {
	db, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error on get DB instance %w", err)
	}

	if err = db.Close(); err != nil {
		return fmt.Errorf("error on close DB instance %w", err)
	}

	return nil
}

func GetInstance() *gorm.DB {
	return instance
}
