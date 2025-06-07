package main

import (
	"log"
	"os"

	"github.com/glendecado/template/migration"
	"github.com/glendecado/template/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//app
	app := fiber.New()

	app.Use(logger.New())

	//route
	route.Run(app)

	// Connect to database (change as needed)
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")+".db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Run auto migrations
	if err := autoMigrateModels(db); err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}
	log.Println("Auto migration completed successfully")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port if not specified
	}

	app.Listen(":" + port)
}

// AutoMigrateModels runs GORM's AutoMigrate on all models in migrate slice
func autoMigrateModels(db *gorm.DB) error {
	for _, model := range migration.Migrate {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}
