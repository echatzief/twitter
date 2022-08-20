package api

import (
	"backend/models"
	"fmt"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateUser(user models.User) []*ValidationError {
	var errors []*ValidationError
	var validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			var field = err.StructNamespace()
			switch field {
			case "User.Email":
				element.Field = "email"
				element.Error = "Please provide a valid email"
			case "User.FirstName":
				element.Field = "firstName"
				element.Error = "Please provide a valid first name"
			case "User.LastName":
				element.Field = "lastName"
				element.Error = "Please provide a valid last name"
			default:
				element.Field = "generic"
				element.Error = "Something went wrong"
			}
			errors = append(errors, &element)
		}
	}
	return errors
}

// GET /users
func GetUsers(c *fiber.Ctx) error {
	return nil
}

// POST /users
func CreateUser(c *fiber.Ctx) error {

	// validate body
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validate user fields
	errors := ValidateUser(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// open database
	currentPath, _ := os.Getwd()
	fmt.Println(path.Join(currentPath, "/database/twitter-app.db"))
	db, err := gorm.Open(sqlite.Open(path.Join(currentPath, "/database/twitter-app.db")), &gorm.Config{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Error",
		})
	}

	user = &models.User{Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}

	db.Create(user)

	return c.JSON(user)
}

// GET /user
func GetUser(c *fiber.Ctx) error {
	return nil
}

// PUT /user/:id
func UpdateUser(c *fiber.Ctx) error {
	return nil
}

// DELETE /user/:id
func DeleteUser(c *fiber.Ctx) error {
	return nil
}
