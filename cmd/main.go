package main

import (
	"log"
	"loyalty_central/internal/database"
	"loyalty_central/internal/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := gorm.Open(postgres.Open("host="+os.Getenv("POSTGRES_HOST")+" user="+os.Getenv("POSTGRES_USER")+" password="+os.Getenv("POSTGRES_PASSWORD")+" dbname="+os.Getenv("POSTGRES_DB")+" port="+os.Getenv("POSTGRES_PORT")+" sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Successfully connected to database")

	migrationErr := database.MigrateDB(db)
	if migrationErr != nil {
		log.Fatalf("Error running migration: %v", migrationErr)
	}
	log.Println("Successfully running database migration")

	app := fiber.New()
	router.Routing(app, db)
	log.Fatal(app.Listen(":3000"))
}
