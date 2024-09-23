package service

import (
	"fmt"
	"time"

	"github.com/Raditsoic/anime-go/database/repository"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
)

type ReviewService struct {
	ReviewRepo repository.ReviewRepo
}

func NewReviewService(repo repository.ReviewRepo) *ReviewService {
	return &ReviewService{ReviewRepo: repo}
}

func (s *ReviewService) SaveReview(input model.NewReview, userID string) (*model.Review, error) {
	var review *model.Review
	var err error

	if input.ID == nil {
		review = &model.Review{
			ID:          "review_" + utils.GenerateUUID(),
			UserID:      userID,
			ContentID:   *input.ContentID,
			ContentType: input.ContentType,
			Rating:      input.Rating,
			Comment:     input.Comment,
			CreatedAt:   time.Now().Format(time.RFC1123),
		}

		err = s.ReviewRepo.Create(review)
		if err != nil {
			return nil, err
		}
	} else {
		review, err = s.ReviewRepo.GetReviewByID(*input.ID)
		if err != nil {
			return nil, err
		}

		review.Comment = input.Comment
		review.Rating = input.Rating

		err = s.ReviewRepo.UpdateReview(review)
		if err != nil {
			return nil, fmt.Errorf("failed to update review: %v", err)
		}
	}

	return review, nil
}
