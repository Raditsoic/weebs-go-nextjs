package repository

import (
	"context"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnimeList []*model.Anime

type AnimeRepo struct {
	client *mongo.Client
}

func NewAnimeRepo(client *mongo.Client) *AnimeRepo {
	return &AnimeRepo{client: client}
}

func (db *AnimeRepo) GetByID(id string) (interface{}, error) {
	col := db.client.Database("weebs").Collection("anime")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var anime model.Anime
	err := col.FindOne(ctx, filter).Decode(&anime)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &anime, nil
}

func (db *AnimeRepo) GetAll() ([]interface{}, error) {
	col := db.client.Database("weebs").Collection("anime")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var animes AnimeList
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &animes); err != nil {
		return nil, err
	}

	var result []interface{}
	for _, anime := range animes {
		result = append(result, anime)
	}

	return result, nil
}

func (db *AnimeRepo) Create(anime *model.Anime) error {
	col := db.client.Database("weebs").Collection("anime")
	_, err := col.InsertOne(context.TODO(), anime)
	return err
}
