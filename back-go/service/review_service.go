package service

import (
	"fmt"
	"time"

	"github.com/Raditsoic/anime-go/database/repository"
	"github.com/Raditsoic/anime-go/graph/model"
)

type ReviewService struct {
	ReviewRepo repository.ReviewRepo
	AnimeRepo  repository.AnimeRepo
	MangaRepo  repository.MangaRepo
}

func NewReviewService(repo repository.ReviewRepo, animeRepo repository.AnimeRepo, mangaRepo repository.MangaRepo) *ReviewService {
	return &ReviewService{
		ReviewRepo: repo,
		AnimeRepo:  animeRepo,
		MangaRepo:  mangaRepo}
}

func (s *ReviewService) SaveReview(input model.NewReview, userID string) (*model.Review, error) {
	if input.ContentType == "anime" {
		_, err := s.AnimeRepo.GetByID(input.ContentID)
		if err != nil {
			return nil, fmt.Errorf("failed to get anime: %v", err)
		}
	} else if input.ContentType == "manga" {
		_, err := s.ReviewRepo.GetReviewByID(input.ContentID)
		if err != nil {
			return nil, fmt.Errorf("failed to get manga: %v", err)
		}
	} else {
		return nil, fmt.Errorf("invalid content type")
	}

	review := &model.Review{
		ContentID:   input.ContentID,
		ContentType: input.ContentType,
		Rating:      input.Rating,
		Comment:     input.Comment,
		CreatedAt:   time.Now().Format(time.RFC1123),
	}
	err := s.ReviewRepo.Create(review)
	if err != nil {
		return nil, err
	}

	return review, nil
}
