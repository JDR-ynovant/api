package webpush

import (
	"errors"
	"github.com/JDR-ynovant/api/internal"
	"github.com/JDR-ynovant/api/internal/models"
	"github.com/JDR-ynovant/api/internal/repository"
	"github.com/SherClockHolmes/webpush-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
)

func SendNotificationToGame(g *models.Game, message string) error {
	playerIds := make([]primitive.ObjectID, 0)
	for _, player := range g.Players {
		playerIds = append(playerIds, player.Id)
	}

	ur := repository.NewUserRepository()
	users, err := ur.FindAllBy(bson.M{"_id": playerIds})

	if err != nil {
		return err
	}

	for _, user := range users {
		err = SendNotificationToPlayer(user, message)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func SendNotificationToPlayers(playersIds []primitive.ObjectID, message string) error {
	ur := repository.NewUserRepository()
	users, err := ur.FindAllBy(bson.M{"_id": playersIds})

	if err != nil {
		return err
	}

	for _, user := range users {
		err = SendNotificationToPlayer(user, message)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func SendNotificationToPlayer(user models.User, message string) error {
	if user.Subscription.Endpoint == "" {
		return errors.New("missing subscription")
	}

	_, err := SendNotification(&user.Subscription, message)
	return err
}

func SendNotification(subscription *webpush.Subscription, message string) (*http.Response, error) {
	config := internal.GetConfig()

	resp, err := webpush.SendNotification([]byte(message), subscription, &webpush.Options{
		Subscriber:      "g.marmo@hotmail.fr",
		VAPIDPublicKey:  config.VapidPublicKey,
		VAPIDPrivateKey: config.VapidPrivateKey,
		TTL:             30,
	})

	if err == nil {
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
	}

	return resp, err
}
