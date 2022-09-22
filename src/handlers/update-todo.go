package handlers

import (
	"fmt"
	"notes-service/src/models"
	"notes-service/src/utils"

	"github.com/gofiber/fiber/v2"
)

func UpdateNote(c *fiber.Ctx) error {
	userID := fmt.Sprintf("%s", c.Locals("uuid"))
	noteID := c.Params("noteID")

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

	// update note
	err = models.UpdateNote(body, noteID, userID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}
