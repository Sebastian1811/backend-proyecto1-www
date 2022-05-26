package config

import (
	"os"

	"github.com/Sebastian1811/backend-proyecto1-www/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConn() *gorm.DB {
	errenv := godotenv.Load()

	if errenv != nil {
		panic("failed load env")
	}
	dburi := os.Getenv("DBURI")

	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}
	db.AutoMigrate(&entity.Beca{}, &entity.User{}, &entity.Requisito{})
	return db
}

func CloseDB(db *gorm.DB) {
	sqldb, _ := db.DB()
	sqldb.Close()
}
