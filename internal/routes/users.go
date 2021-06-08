package routes

import (
	"encoding/json"
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserRouteHandler struct{}

func (UserRouteHandler) Register(app fiber.Router) {
	usersApi := app.Group("/users")

	usersApi.Get("/:id?", getUser)
	usersApi.Post("", createUser)
	usersApi.Put("/:id", updateUser)
	usersApi.Delete("/:id", deleteUser)

	log.Println("Registered users api group.")
}

// getUser godoc
// @Summary Show a user
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func getUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	if c.Params("id") == "" {
		users, err := ur.FindAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(users)
	}

	user, err := ur.FindOneById(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(user)
}

// createUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /users [post]
func createUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	user := models.User{
		Games: make([]primitive.ObjectID, 0),
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	u, err := ur.Create(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(*u)
}

// updateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param id path int true "User ID"
// @Router /users/{id} [put]
func updateUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	var user models.User
	json.Unmarshal(c.Body(), &user)

	err := ur.Update(c.Params("id"), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(user)
}

// deleteUser godoc
// @Summary Delete an existing user
// @Description Delete an existing user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param id path int true "User ID"
// @Router /users/{id} [delete]
func deleteUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	err := ur.Delete(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
