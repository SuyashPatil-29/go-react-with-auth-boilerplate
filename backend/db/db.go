package db

import (
	"backend/internal/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Failed to load .env file:", err)
	}

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	log.Info().Msg("Database connected successfully")

	// migrate the schema
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Notebook{},
		&models.Chapter{},
		&models.Notes{},
	); err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate schema")
	}
}
