package repository

import (
	"context"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListAnime *model.AnimeList
type AnimeListEntry *model.AnimeListEntries
type AnimeListRepo struct {
	client *mongo.Client
}

func NewAnimeListRepo(client *mongo.Client) *AnimeListRepo {
	return &AnimeListRepo{client: client}
}

func (db *AnimeListRepo) GetByID(id string) (interface{}, error) {
	col := db.client.Database("weebs").Collection("anime_list")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	var animeList model.AnimeList
	err := col.FindOne(ctx, filter).Decode(&animeList)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &animeList, nil
}

func (db *AnimeListRepo) GetAnimeListByUserID(user_id string) (interface{}, error) {
	col := db.client.Database("weebs").Collection("anime_list")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"userID": user_id}
	var animeList model.AnimeList
	err := col.FindOne(ctx, filter).Decode(&animeList)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &animeList, nil
}

func (db *AnimeListRepo) UpdateAnimeList(AnimeList *model.AnimeList) error {
	col := db.client.Database("weebs").Collection("anime_list")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"userID": AnimeList.UserID}
	update := bson.M{
		"$set": bson.M{
			"entries": AnimeList.Entries,
		},
	}

	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
