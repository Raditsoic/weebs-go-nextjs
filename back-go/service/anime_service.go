package service

import (
	"fmt"

	"github.com/Raditsoic/anime-go/database"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
)

type AnimeService struct {
	AnimeRepo database.AnimeRepo
}

func NewAnimeService(repo database.AnimeRepo) *AnimeService {
	return &AnimeService{AnimeRepo: repo}
}

func (s *AnimeService) CreateAnime(input model.NewAnime) (*model.Anime, error) {
	anime := &model.Anime{
		ID:          fmt.Sprint(utils.GenerateUUID()),
		Title:       input.Title,
		Image:       input.Image,
		Description: input.Description,
		Genre: &model.Genre{
			ID:   input.GenreID,
			Name: "Genre " + input.GenreID,
		},
	}
	if err := s.AnimeRepo.Create(anime); err != nil {
		return nil, err
	}
	return anime, nil
}

func (s *AnimeService) GetAnimes() ([]*model.Anime, error) {
	result, err := s.AnimeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	animeList := make([]*model.Anime, len(result))
	for i, item := range result {
		anime, ok := item.(*model.Anime)
		if !ok {
			return nil, fmt.Errorf("unexpected type in AnimeRepo: %T", item)
		}
		animeList[i] = anime
	}

	return animeList, nil
}

func (s *AnimeService) GetAnimeByID(id string) (*model.Anime, error) {
	animeInterface, err := s.AnimeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	anime, ok := animeInterface.(*model.Anime)
	if !ok {
		return nil, fmt.Errorf("unexpected type in AnimeRepo: %T", animeInterface)
	}
	return anime, nil
}
