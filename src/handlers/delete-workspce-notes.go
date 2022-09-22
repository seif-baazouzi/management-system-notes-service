package handlers

import (
	"notes-service/src/models"
	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteWorkspaceNotes(c *fiber.Ctx) error {
	workspaceID := c.Params("workspaceID")

	err := models.DeleteWorkspaceNotes(workspaceID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
