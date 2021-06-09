package repository

import (
	"context"
	"errors"
	"github.com/JDR-ynovant/api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type GameRepository struct {
	CollectionName string
	Collection     *mongo.Collection
	Context        context.Context
}

func NewGameRepository() GameRepository {
	collectionName := "Games"
	c, err := GetMongoDbCollection(collectionName)

	if err != nil {
		log.Fatalln(err)
	}

	return GameRepository{
		CollectionName: collectionName,
		Collection:     c,
		Context:        context.Background(),
	}
}

func (gr GameRepository) FindAll() ([]models.Game, error) {
	if gr.Collection == nil {
		return nil, errors.New("missing connection")
	}

	var results []models.Game
	cur, err := gr.Collection.Find(gr.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	cur.All(gr.Context, &results)

	return results, err
}

func (gr GameRepository) FindAllBy(filter bson.M) ([]models.Game, error) {
	var fetchedGame []models.Game

	cur, err := gr.Collection.Find(gr.Context, filter)
	if err != nil {
		return nil, err
	}

	cur.All(gr.Context, &fetchedGame)

	return fetchedGame, nil
}

func (gr GameRepository) FindOneById(id string) (*models.Game, error) {
	if gr.Collection == nil {
		return nil, errors.New("missing connection")
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	var fetchedGame models.Game

	err := gr.Collection.FindOne(gr.Context, filter).Decode(&fetchedGame)
	if err != nil {
		return nil, err
	}

	return &fetchedGame, nil
}

func (gr GameRepository) Create(u *models.Game) (*models.Game, error) {
	if gr.Collection == nil {
		return nil, errors.New("missing connection")
	}

	u.Id = primitive.NewObjectID()
	_, err := gr.Collection.InsertOne(gr.Context, *u)

	return u, err
}

func (gr GameRepository) Update(id string, u models.Game) error {
	if gr.Collection == nil {
		return errors.New("missing connection")
	}

	update := bson.M{
		"$set": u,
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := gr.Collection.UpdateOne(gr.Context, bson.M{"_id": objID}, update)

	return err
}

func (gr GameRepository) Delete(id string) error {
	if gr.Collection == nil {
		return errors.New("missing connection")
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := gr.Collection.DeleteOne(gr.Context, bson.M{"_id": objID})

	return err
}

func (gr GameRepository) AttachUser(userID string, gameID string) error {
	ur := NewUserRepository()
	user, err := ur.FindOneById(userID)
	if err != nil {
		return err
	}

	if !user.HasGameAttached(gameID) {
		gameIDPrimitive, _ := primitive.ObjectIDFromHex(gameID)
		user.Games = append(user.Games, gameIDPrimitive)

		err = ur.Update(userID, *user)
	}

	return err
}