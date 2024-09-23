package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrReviewNotFound = errors.New("Review not found.")

// ReviewRepo represents the database operations for reviews.
type ReviewRepo struct {
	client *mongo.Client
}

// NewReviewRepo creates a new ReviewRepo instance.
func NewReviewRepo(client *mongo.Client) *ReviewRepo {
	return &ReviewRepo{client: client}
}

// Create inserts a new review into the database.
func (db *ReviewRepo) Create(review *model.Review) error {
	col := db.client.Database("weebs").Collection("reviews")
	_, err := col.InsertOne(context.TODO(), review)
	return err
}

// GetReviewByID retrieves a review by ID.
func (db *ReviewRepo) GetReviewByID(id string) (*model.Review, error) {
	col := db.client.Database("weebs").Collection("reviews")
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	filter := bson.M{"_id": objectID}
	var review model.Review

	err = col.FindOne(ctx, filter).Decode(&review)
	if err == mongo.ErrNoDocuments {
		return nil, ErrReviewNotFound
	}

	if err != nil {
		return nil, err
	}

	return &review, nil
}

// UpdateReview updates an existing review in the database.
func (db *ReviewRepo) UpdateReview(review *model.Review) error {
	col := db.client.Database("weebs").Collection("reviews")
	ctx, cancel := utils.GetContextWithTimeout()
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(review.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": review,
	}

	_, err = col.UpdateOne(ctx, filter, update)
	return err
}