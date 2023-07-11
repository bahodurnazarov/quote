package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"quote/internal/models"
	"quote/pkg/db"
)

type User struct {
	// this is not the model User, see this as the serialzer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreteResponseUser(userModer models.User) User {
	return User{ID: userModer.ID, FirstName: userModer.FirstName, LastName: userModer.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Database.DB.Create(&user)
	responseUser := CreteResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	db.Database.DB.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreteResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	db.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreteResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updataData UpdateUser

	if err := c.BodyParser(&updataData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updataData.FirstName
	user.LastName = updataData.LastName

	db.Database.DB.Save(&user)

	responseUser := CreteResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := db.Database.DB.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully Deleted User")
}
