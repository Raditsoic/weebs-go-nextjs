package database

import (
	"context"
	"log"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/Raditsoic/anime-go/utils"
)

var connectionString string = "mongodb://localhost:27017/"
const databaseName string = "weebs"
const animeCollection string = "anime"
const mangaCollection string = "manga"
const userCollection string = "users"

const maxPoolSize = 50
const timeoutSeconds = 30

type AnimeList []*model.Anime
type MangaList []*model.Manga

type RepoInterface interface {
	GetByID(id string) (interface{}, error)
	GetAll() ([]interface{}, error)
	Create()
}

// AnimeRepo represents the database operations for anime.
type AnimeRepo struct {
	client *mongo.Client
}

// MangaRepo represents the database operations for manga.
type MangaRepo struct {
	client *mongo.Client
}

type UserRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) *UserRepo {
	return &UserRepo{client: client}
}

// NewAnimeRepo creates a new AnimeRepo instance.
func NewAnimeRepo(client *mongo.Client) *AnimeRepo {
	return &AnimeRepo{client: client}
}

// NewMangaRepo creates a new MangaRepo instance.
func NewMangaRepo(client *mongo.Client) *MangaRepo {
	return &MangaRepo{client: client}
}

// Connect establishes a connection to the MongoRepo database.
func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions = clientOptions.SetMaxPoolSize(maxPoolSize)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// GetByID retrieves a record from the database based on the provided ID.
func (db *AnimeRepo) GetByID(id string) (interface{}, error) {
	col := db.client.Database(databaseName).Collection(animeCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	_id := id
	filter := bson.M{"_id": _id}
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

// GetByID retrieves a record from the database based on the provided ID.
func (db *MangaRepo) GetByID(id string) (interface{}, error) {
	col := db.client.Database(databaseName).Collection(mangaCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
	defer cancel()

	_id := id
	filter := bson.M{"_id": _id}
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

// GetAll retrieves all records from the database.
func (db *AnimeRepo) GetAll() ([]interface{}, error) {
	col := db.client.Database(databaseName).Collection(animeCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
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

// GetAll retrieves all records from the database.
func (db *MangaRepo) GetAll() ([]interface{}, error) {
	col := db.client.Database(databaseName).Collection(mangaCollection)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
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

func (db *AnimeRepo) Create(anime *model.Anime) error {
	animeCol := db.client.Database(databaseName).Collection(animeCollection)
	_, err := animeCol.InsertOne(context.TODO(), anime)
	if err != nil {
		return err
	}
	return nil
}

func (db *MangaRepo) Create(manga *model.Manga) error {
	mangaCol := db.client.Database(databaseName).Collection(mangaCollection)
	_, err := mangaCol.InsertOne(context.TODO(), manga)
	if err != nil {
		return err
	}
	return nil
}


func (db *UserRepo) Register(user *model.User) error {
	userCol := db.client.Database(databaseName).Collection(userCollection)
	_, err := userCol.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (db *UserRepo) GetByEmail(email string) (*model.User, error) {
	col := db.client.Database(databaseName).Collection(userCollection)
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	filter := bson.M{"email": email}
	var user model.User
	err := col.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *UserRepo) GetByUsername(username string) (*model.User, error) {
	col := db.client.Database(databaseName).Collection(userCollection)
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	filter := bson.M{"username": username}
	var user model.User
	err := col.FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

