package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"quote/internal/models"
	"quote/pkg/db"
	"time"
)

//{
//	id: 1,
//	user: {
//		id: 23,
//		first_name: "Bob marle",
//		last_name: "Marley"
//	},
//	product : {
//		id :24,
//		name: "Mackbook",
//		serial_number: "23423"
//	}
//}

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.UserRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Database.DB.Create(&order)

	responseUser := CreteResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	db.Database.DB.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		db.Database.DB.Find(&user, "id = ?", order.UserRefer)
		db.Database.DB.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreteResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	db.Database.DB.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}

	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	db.Database.DB.First(&user, order.UserRefer)
	db.Database.DB.First(&product, order.ProductRefer)
	responseUser := CreteResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}
