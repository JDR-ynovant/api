package routes

import (
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"log"
)

type UserRouteHandler struct{}

func (UserRouteHandler) Register(app fiber.Router) {
	usersApi := app.Group("/users")

	usersApi.Get("/", getUsers)
	usersApi.Get("/:id", getUser)
	usersApi.Post("", createUser)
	usersApi.Put("/:id", updateUser)
	usersApi.Delete("/:id", deleteUser)

	log.Println("Registered users api group.")
}

// getUsers godoc
// @Summary List users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 array []models.User
// @Router /users [get]
func getUsers(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()
	users, err := ur.FindAll()
	if err != nil {
		return jsonError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(users)
}

// getUser godoc
// @Summary Show a user
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func getUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	user, err := ur.FindOneById(c.Params("id"))
	if err != nil {
		return jsonError(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}

type CreateUserRequest struct {
	Name string `validate:"required"`
}

func ValidateCreateUserRequest(userRequest CreateUserRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(userRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
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

	createUserRequest := new(CreateUserRequest)
	if err := c.BodyParser(createUserRequest); err != nil {
		return jsonError(c, fiber.StatusInternalServerError, err.Error())
	}

	validationErrors := ValidateCreateUserRequest(*createUserRequest)
	if validationErrors != nil {
		return jsonError(c, fiber.StatusBadRequest, validationErrors)
	}

	user := models.User{
		Name:  createUserRequest.Name,
		Games: make([]models.MetaGame, 0),
	}

	u, err := ur.Create(&user)
	if err != nil {
		return jsonError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(*u)
}

type UpdateUserRequest struct {
	Name string `validate:"required"`
}

func ValidateUpdateUserRequest(userRequest UpdateUserRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(userRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// updateUser godoc
// @Summary Update an existing user
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param id path string true "User ID"
// @Router /users/{id} [put]
func updateUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	updateUserRequest := new(UpdateUserRequest)
	if err := c.BodyParser(updateUserRequest); err != nil {
		return jsonError(c, fiber.StatusInternalServerError, err.Error())
	}

	validationErrors := ValidateUpdateUserRequest(*updateUserRequest)
	if validationErrors != nil {
		return jsonError(c, fiber.StatusBadRequest, validationErrors)
	}

	user, err := ur.FindOneById(c.Params("id"))
	if err != nil {
		return jsonError(c, fiber.StatusNotFound, err.Error())
	}

	user.Name = updateUserRequest.Name
	err = ur.Update(c.Params("id"), *user)
	if err != nil {
		return jsonError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(updateUserRequest)
}

// deleteUser godoc
// @Summary Delete an existing user
// @Description Delete an existing user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
func deleteUser(c *fiber.Ctx) error {
	ur := repository.NewUserRepository()

	err := ur.Delete(c.Params("id"))
	if err != nil {
		return jsonError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
