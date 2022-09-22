package handlers

import (
	"notes-service/src/models"
	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceID")

	notes := []models.Note{}
	notes, err := models.GetNotes(workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"notes": notes})
}
