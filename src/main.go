package main

import (
	"fmt"
	"notes-service/src/auth"
	"notes-service/src/db"
	"notes-service/src/handlers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.InitDataBase()
	defer db.CloseDataBase()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/api/v1/notes/:workspaceID", auth.IsWorkspaceOwner, handlers.GetNotes)
	app.Get("/api/v1/notes/single/:noteID", auth.IsLogin, handlers.GetSingleNote)
	app.Post("/api/v1/notes/:workspaceID", auth.IsWorkspaceOwner, handlers.CreateNote)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
