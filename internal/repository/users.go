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

type UserRepository struct {
	CollectionName string
	Collection     *mongo.Collection
	Context        context.Context
}

func NewUserRepository() UserRepository {
	collectionName := "Users"
	c, err := GetMongoDbCollection(collectionName)

	if err != nil {
		log.Fatalln(err)
	}

	return UserRepository{
		CollectionName: collectionName,
		Collection:     c,
		Context:        context.Background(),
	}
}

func (ur UserRepository) FindAll() ([]models.User, error) {
	if ur.Collection == nil {
		return nil, errors.New("missing connection")
	}

	var results []models.User
	cur, err := ur.Collection.Find(ur.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	cur.All(ur.Context, &results)

	return results, err
}

func (ur UserRepository) FindAllBy(filter bson.M) ([]models.User, error) {
	var fetchedUser []models.User

	cur, err := ur.Collection.Find(ur.Context, filter)
	if err != nil {
		return nil, err
	}

	cur.All(ur.Context, &fetchedUser)

	return fetchedUser, nil
}

func (ur UserRepository) FindOneById(id string) (*models.User, error) {
	if ur.Collection == nil {
		return nil, errors.New("missing connection")
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	var fetchedUser models.User

	err := ur.Collection.FindOne(ur.Context, filter).Decode(&fetchedUser)
	if err != nil {
		return nil, err
	}

	return &fetchedUser, nil
}

func (ur UserRepository) Create(u *models.User) (*models.User, error) {
	if ur.Collection == nil {
		return nil, errors.New("missing connection")
	}

	u.Id = primitive.NewObjectID()
	_, err := ur.Collection.InsertOne(ur.Context, u)

	return u, err
}

func (ur UserRepository) Update(id string, u models.User) error {
	if ur.Collection == nil {
		return errors.New("missing connection")
	}

	update := bson.M{
		"$set": u,
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ur.Collection.UpdateOne(ur.Context, bson.M{"_id": objID}, update)

	return err
}

func (ur UserRepository) Delete(id string) error {
	if ur.Collection == nil {
		return errors.New("missing connection")
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := ur.Collection.DeleteOne(ur.Context, bson.M{"_id": objID})

	return err
}
