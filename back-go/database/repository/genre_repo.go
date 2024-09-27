package repository

import (
	"context"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Genre []*model.Genre

type GenreRepo struct {
	client *mongo.Client
}

func NewGenreRepo(client *mongo.Client) *GenreRepo {
	return &GenreRepo{client: client}
}

func (db *GenreRepo) GetByID(id string) (interface{}, error) {
	col := db.client.Database("weebs").Collection("genre")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	var genre model.Genre
	err := col.FindOne(ctx, filter).Decode(&genre)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (db *GenreRepo) Create(genre *model.Genre) error {
	col := db.client.Database("weebs").Collection("genre")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := col.InsertOne(ctx, genre)
	if err != nil {
		return err
	}

	return nil
}
