package main

import (
	"log"

	"github.com/joho/godotenv"
	deliverhttp "gitlab.com/rizkyimaduddin24/techtest/internal/delivery/http"
	pgdb "gitlab.com/rizkyimaduddin24/techtest/internal/infrastructure/postgres"
)

func init() {
	// Load Environtment
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	// Connect to PostgresDB
	if err := pgdb.ConnectToDB(); err != nil {
		log.Fatal(err.Error())
	}

	// Auto Migrate User Table
	// migrate.User()
}

func main() {
	// Serve HTTP
	deliverhttp.Serve(pgdb.DB)
}
