package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"quote/internal/controllers"
)

func Route() {
	//db.DBconn()

	app := fiber.New()
	setupRoutes(app)

	log.Println("Hello! kdfg" + "sdgdfg")
	fmt.Println("Worlddf")
	log.Println("Hello!")

	port := os.Getenv("S_PORT")
	log.Fatal(app.Listen(port))
}

func setupRoutes(app *fiber.App) {
	// User endpoints
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetUsers)
	log.Println("Hello!")
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)
	// Product endpoints
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products", controllers.GetProducts)
	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)
	// Order endpoints
	app.Post("/api/orders", controllers.CreateOrder)
	app.Get("/api/orders", controllers.GetOrders)
	app.Get("/api/orders/:id", controllers.GetOrder)
}
