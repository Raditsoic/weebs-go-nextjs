package repository

import (
	"context"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ListManga *model.MangaList
type MangaListEntry *model.MangaListEntries

type MangaListRepo struct {
	client *mongo.Client
}

func NewMangaListRepo(client *mongo.Client) *MangaListRepo {
	return &MangaListRepo{client: client}
}

func (db *MangaListRepo) GetMangaListByUserID(user_id string) (interface{}, error) {
	col := db.client.Database("weebs").Collection("manga_list")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"userID": user_id}
	var mangaList model.MangaList
	err := col.FindOne(ctx, filter).Decode(&mangaList)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &mangaList, nil
}

func (db *MangaListRepo) UpdateMangaList(MangaList *model.MangaList) error {
	col := db.client.Database("weebs").Collection("manga_list")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"userID": MangaList.UserID}
	update := bson.M{
		"$set": bson.M{
			"entries": MangaList.Entries,
		},
	}

	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
