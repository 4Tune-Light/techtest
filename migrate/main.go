package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/rizkyimaduddin24/techtest/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// Load Environtment
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	// Load Env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	if dbHost == "" || dbUser == "" || dbPass == "" || dbName == "" || dbPort == "" {
		log.Fatal("database configuration invalid")
	}

	// Connect DB
	var err error
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=UTC"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	// Check table for `User` exists or not
	DB.Migrator().HasTable(&entity.User{})

	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	DB.Migrator().DropTable(&entity.User{})

	// Run migration
	if err := DB.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal("migration failed")
	}

	// Run seeder
	DB.Create(&entity.User{
		Name:     "admin",
		Email:    "admin@email.com",
		Role:     "admin",
		Password: "$2a$10$1T.cIPGMd44VBeKe8wXejuKbb0bwLogG9yCEkq1GwJJ51VQuYzCaa",
	})
	DB.Create(&entity.User{
		Name:     "user1",
		Email:    "user1@email.com",
		Role:     "user1",
		Password: "$2a$10$1T.cIPGMd44VBeKe8wXejuKbb0bwLogG9yCEkq1GwJJ51VQuYzCaa",
	})
	DB.Create(&entity.User{
		Name:     "user2",
		Email:    "user2@email.com",
		Role:     "user2",
		Password: "$2a$10$1T.cIPGMd44VBeKe8wXejuKbb0bwLogG9yCEkq1GwJJ51VQuYzCaa",
	})
	DB.Create(&entity.User{
		Name:     "user3",
		Email:    "user3@email.com",
		Role:     "user3",
		Password: "$2a$10$1T.cIPGMd44VBeKe8wXejuKbb0bwLogG9yCEkq1GwJJ51VQuYzCaa",
	})
}
