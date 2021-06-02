package routes

import (
	"context"
	"encoding/json"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserRouteHandler struct{}

func (UserRouteHandler) Register(app *fiber.App) {
	app.Get("/users/:id?", getUser)
	app.Post("/users", createUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)
	log.Println("Registered users endpoint routes.")
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
	return nil
}

func updateUser(c *fiber.Ctx) error {
	return nil
}

func deleteUser(c *fiber.Ctx) error {
	return nil
}