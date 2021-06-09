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

type GridRepository struct {
	CollectionName string
	Collection     *mongo.Collection
	Context        context.Context
}

func NewGridRepository() GridRepository {
	collectionName := "Grids"
	c, err := GetMongoDbCollection(collectionName)

	if err != nil {
		log.Fatalln(err)
	}

	return GridRepository{
		CollectionName: collectionName,
		Collection:     c,
		Context:        context.Background(),
	}
}

func (gr GridRepository) FindAll() ([]models.Grid, error) {
	if gr.Collection == nil {
		return nil, errors.New("missing connection")
	}

	var results []models.Grid
	cur, err := gr.Collection.Find(gr.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	cur.All(gr.Context, &results)

	return results, err
}

func (gr GridRepository) FindAllBy(filter bson.M) ([]models.Grid, error) {
	var fetchedGrid []models.Grid

	cur, err := gr.Collection.Find(gr.Context, filter)
	if err != nil {
		return nil, err
	}

	cur.All(gr.Context, &fetchedGrid)

	return fetchedGrid, nil
}

func (gr GridRepository) FindOneById(id string) (*models.Grid, error) {
	if gr.Collection == nil {
		return nil, errors.New("missing connection")
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	var fetchedGrid models.Grid

	err := gr.Collection.FindOne(gr.Context, filter).Decode(&fetchedGrid)
	if err != nil {
		return nil, err
	}

	return &fetchedGrid, nil
}

func (gr GridRepository) Create(u *models.Grid) error {
	if gr.Collection == nil {
		return errors.New("missing connection")
	}

	u.Id = primitive.NewObjectID()
	_, err := gr.Collection.InsertOne(gr.Context, *u)

	return err
}

func (gr GridRepository) Update(id string, u models.Grid) error {
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

func (gr GridRepository) Delete(id string) error {
	if gr.Collection == nil {
		return errors.New("missing connection")
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := gr.Collection.DeleteOne(gr.Context, bson.M{"_id": objID})

	return err
}
