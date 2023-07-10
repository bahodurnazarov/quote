package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"quote/pkg/db"
)

func Route() {
	app := fiber.New()
	db.Postgres()
	port := os.Getenv("S_PORT")

	//app.Get("/", waiting)
	// books.RegisterRoutes(app, db)
	log.Println(port)
	app.Listen(port)
}
