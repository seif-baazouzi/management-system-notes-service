package handlers

import (
	"fmt"
	"notes-service/src/models"
	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func GetSingleNote(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	noteID := c.Params("noteID")

	note := models.Note{}
	found, note, err := models.GetSingleNote(noteID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !found {
		return c.Status(404).JSON(fiber.Map{"note": nil})
	}

	return c.JSON(fiber.Map{"note": note})
}
