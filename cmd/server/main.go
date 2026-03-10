package main

import (
	"MathTrainer/internal/database"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Info("unable to load .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		logger.Error("no connection string")
		return
	}

	db, err := database.OpenDB(connectionString)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
	}
	defer db.Close()

	// server = http.Server{}
}
