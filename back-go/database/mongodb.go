<<<<<<< HEAD
package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxPoolSize = 50
const timeoutSeconds = 30

var client *mongo.Client

func Connect() *mongo.Client {
	if client == nil {
		// Check if MONGO_URI is set, otherwise use the default connectionString
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			mongoURI = "mongodb://mongo:27017"
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		clientOptions = clientOptions.SetMaxPoolSize(maxPoolSize)

		ctx, cancel := context.WithTimeout(context.Background(), timeoutSeconds*time.Second)
		defer cancel()

		var err error
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		// Ping the database to verify connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}

		log.Println("Connected to MongoDB at", mongoURI)
	}

	return client
}
=======
package database

import (
	"context"
	"log"
	"time"

	"github.com/Raditsoic/anime-go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb://localhost:27017/"

const databaseName string = "weebs"
const animeCollection string = "anime"
const mangaCollection string = "manga"

const maxPoolSize = 50
const timeoutSeconds = 30

type AnimeList []*model.Anime
type MangaList []*model.Manga

type DBInterface interface {
	GetByID(id string) (interface{}, error)
	GetAll() ([]interface{}, error)
	Create()
}

// AnimeDB represents the database operations for anime.
type AnimeDB struct {
	client *mongo.Client
}

// MangaDB represents the database operations for manga.
type MangaDB struct {
	client *mongo.Client
}

// NewAnimeDB creates a new AnimeDB instance.
func NewAnimeDB(client *mongo.Client) *AnimeDB {
	return &AnimeDB{client: client}
}

// NewMangaDB creates a new MangaDB instance.
func NewMangaDB(client *mongo.Client) *MangaDB {
	return &MangaDB{client: client}
}

// Connect establishes a connection to the MongoDB database.
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
func (db *AnimeDB) GetByID(id string) (interface{}, error) {
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
func (db *MangaDB) GetByID(id string) (interface{}, error) {
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
func (db *AnimeDB) GetAll() ([]interface{}, error) {
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
func (db *MangaDB) GetAll() ([]interface{}, error) {
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

func (db *AnimeDB) Create(anime *model.Anime) {
	animeCol := db.client.Database(databaseName).Collection(animeCollection)
	_, err := animeCol.InsertOne(context.TODO(), anime)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *MangaDB) Create(manga *model.Manga) {
	mangaCol := db.client.Database(databaseName).Collection(mangaCollection)
	_, err := mangaCol.InsertOne(context.TODO(), manga)
	if err != nil {
		log.Fatal(err)
	}
}
>>>>>>> 9b9b4dc (Add: front-end)
