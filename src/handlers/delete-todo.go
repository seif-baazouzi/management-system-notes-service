package handlers

import (
	"fmt"
	"notes-service/src/models"
	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteNote(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	noteID := c.Params("noteID")

	err := models.DeleteNote(noteID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
