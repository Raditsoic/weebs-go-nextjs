package repository

import (
	"context"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MangaList []*model.Manga

type MangaRepo struct {
	client *mongo.Client
}

func NewMangaRepo(client *mongo.Client) *MangaRepo {
	return &MangaRepo{client: client}
}

func (db *MangaRepo) GetByID(id string) (interface{}, error) {
	col := db.client.Database("weebs").Collection("manga")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var manga model.Manga
	err := col.FindOne(ctx, filter).Decode(&manga)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &manga, nil
}

func (db *MangaRepo) GetAll() ([]interface{}, error) {
	col := db.client.Database("weebs").Collection("manga")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var mangas MangaList
	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &mangas); err != nil {
		return nil, err
	}

	var result []interface{}
	for _, manga := range mangas {
		result = append(result, manga)
	}

	return result, nil
}

func (db *MangaRepo) Create(manga *model.Manga) error {
	manga.ID = "manga" + utils.GenerateUUID()

	col := db.client.Database("weebs").Collection("manga")
	_, err := col.InsertOne(context.TODO(), manga)
	return err
}
