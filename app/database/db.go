package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

var DB Database

func ConnectDb() (Database, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatalf("DB_URL is not set in the environment variables")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Warn),
	})
	if err != nil {
		return Database{}, err
	}

	DB = Database{Db: db}

	sqlDB, err := DB.Db.DB()
	if err != nil {
		return Database{}, err
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	return DB, nil
}
