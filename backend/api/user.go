package api

import (
	"backend/database"
	"backend/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// validate user fields
func validateUser(user models.User) []*ValidationError {
	var errors []*ValidationError
	var validate = validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			var field = err.StructNamespace()
			switch field {
			case "User.Username":
				element.Field = "username"
				element.Error = "Please provide a valid username"
			case "User.Password":
				element.Field = "password"
				element.Error = "Please provide a valid username"
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
	users := []models.User{}
	database.Database.Handler.Find(&users)
	return c.Status(200).JSON(users)
}

// GET /user/:id
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Please ensure that id is an integer",
		})
	}

	database.Database.Handler.Find(&user, "id = ?", id)
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{})
	}
	return c.JSON(user)
}

// POST /users
func CreateUser(c *fiber.Ctx) error {

	// validate if body is given
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// validate user fields
	errors := validateUser(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// check if user exists with same email or username then throw an error
	existingUser := new(models.User)
	database.Database.Handler.Where(&models.User{}).Find(&existingUser, "email = ? or username = ?", user.Email, user.Username)
	if existingUser.Id != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User exists with given username or email",
		})
	}

	// create user
	database.Database.Handler.Create(&user)

	return c.JSON(user)
}

// PUT /user/:id
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Please ensure that id is an integer",
		})
	}

	// validate if body is given
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check if user that we want to update already exists
	existingUser := new(models.User)
	database.Database.Handler.Where(&models.User{}).Find(&existingUser, "id = ?", id)
	if existingUser.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User that you are trying to update doesn't exist",
		})
	}

	// update user with new fields
	database.Database.Handler.Model(&models.User{}).Where("id = ?", id).Updates(user)

	// find user after update
	user = new(models.User)
	database.Database.Handler.Where(&models.User{}).Find(&user, "id = ?", id)

	return c.JSON(user)
}

// DELETE /user/:id
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Please ensure that id is an integer",
		})
	}

	database.Database.Handler.Find(&user, "id = ?", id)
	if user.Id == 0 {
		return c.Status(404).JSON(fiber.Map{})
	}

	if err = database.Database.Handler.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(user)
}
