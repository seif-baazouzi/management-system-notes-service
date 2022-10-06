package handlers

import (
	"fmt"
	"notes-service/src/models"
	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	workspaceID := c.Params("workspaceID")

	// parse body
	var body models.NoteBody

	err := c.BodyParser(&body)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// checks
	if body.Title == "" {
		return c.JSON((fiber.Map{"title": "Must not be empty"}))
	}

	// create note
	noteID, err := models.CreateNote(body, workspaceID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "noteID": noteID})
}
