package lead

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/leads")
	api.Post("/", CreateLead)
	api.Get("/", GetLeads)
	api.Get("/:id", GetLead)
	api.Put("/:id", UpdateLead)
	api.Delete("/:id", DeleteLead)
}
