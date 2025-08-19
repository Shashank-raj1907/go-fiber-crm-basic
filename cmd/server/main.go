package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/Shashank-raj1907/go-fiber-crm-basic/internal/database"
	"github.com/Shashank-raj1907/go-fiber-crm-basic/internal/lead"
)

func main() {
	// Connect to DB
	database.Connect()
	defer database.Close()

	// Auto-migrate models here
	database.DBConn.AutoMigrate(&lead.Lead{})

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	lead.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
