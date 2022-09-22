package main

import (
	"fmt"
	"notes-service/src/db"
	"os"

	"github.com/gofiber/fiber"
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

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
