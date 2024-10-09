package main

import (
	"github.com/Ashmn07/ticket-booking/config"
	"github.com/Ashmn07/ticket-booking/db"
	"github.com/Ashmn07/ticket-booking/handlers"
	"github.com/Ashmn07/ticket-booking/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
