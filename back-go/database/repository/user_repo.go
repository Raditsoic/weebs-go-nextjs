package repository

import (
	"context"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo represents the database operations for users.
type UserRepo struct {
	client *mongo.Client
}

// NewUserRepo creates a new UserRepo instance.
func NewUserRepo(client *mongo.Client) *UserRepo {
	return &UserRepo{client: client}
}

// Register inserts a new user into the database.
func (db *UserRepo) Register(user *model.User) error {
	col := db.client.Database("weebs").Collection("users")
	_, err := col.InsertOne(context.TODO(), user)
	return err
}

// GetByEmail retrieves a user by email.
func (db *UserRepo) GetByEmail(email string) (*model.User, error) {
	col := db.client.Database("weebs").Collection("users")
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

// GetByUsername retrieves a user by username.
func (db *UserRepo) GetByUsername(username string) (*model.User, error) {
	col := db.client.Database("weebs").Collection("users")
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
