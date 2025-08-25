package lead

import (
	"strconv"

	"github.com/shraj19/go-fiber-crm-basic/internal/database"
	"github.com/gofiber/fiber/v2"
)

func CreateLead(c *fiber.Ctx) error {
	var lead Lead
	if err := c.BodyParser(&lead); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := database.DBConn.Create(&lead).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(lead)
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	if err := database.DBConn.Find(&leads).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var lead Lead
	if err := database.DBConn.First(&lead, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Lead not found"})
	}
	return c.JSON(lead)
}

func UpdateLead(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var lead Lead
	if err := database.DBConn.First(&lead, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Lead not found"})
	}
	if err := c.BodyParser(&lead); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DBConn.Save(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := database.DBConn.Delete(&Lead{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
