package routes

import (
	"fmt"
	"github.com/JDR-ynovant/api/internal/middleware/auth"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"log"
)

type PushRouteHandler struct{}
type SubscriptionRequest struct {
	Subscription webpush.Subscription `validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(subscription SubscriptionRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(subscription)
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

func (PushRouteHandler) Register(app fiber.Router) {
	app.Post("/subscribe", handleSubscribe)

	log.Println("Registered push api group.")
}

// handleSubscribe godoc
// @Summary Subscribe to Push Notification
// @Description Allow to subscribe to game push notification
// @Tags subscribe
// @Accept  json
// @Produce  json
// @Param subscription body SubscriptionRequest true "Subscription data"
// @Success 200
// @Router /api/subscribe [post]
func handleSubscribe(c *fiber.Ctx) error {
	authUser := fmt.Sprintf("%s", c.Locals(auth.ConfigDefault.ContextKey))
	if authUser == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "missing X-User header.",
		})
	}

	ur := repository.NewUserRepository()
	subscription := new(SubscriptionRequest)
	if err := c.BodyParser(subscription); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validationErrors := ValidateStruct(*subscription)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": validationErrors,
		})
	}

	fetchedUser, err := ur.FindOneById(authUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	fetchedUser.Subscription = subscription.Subscription
	err = ur.Update(fetchedUser.Id.Hex(), *fetchedUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
