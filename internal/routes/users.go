package routes

import (
	"context"
	"encoding/json"
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserRouteHandler struct{}

func (UserRouteHandler) Register(app *fiber.App) {
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
	collection, err := repository.GetMongoDbCollection("users")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return err
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return err
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return nil
	}

	json, _ := json.Marshal(results)
	c.Send(json)
	return nil
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
	collection, err := repository.GetMongoDbCollection("users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var person models.User
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response, _ := json.Marshal(res)
	return c.JSON(response)
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
	return nil
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
	return nil
}