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
	note, err := models.GetSingleNote(userID, noteID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"note": note})
}
