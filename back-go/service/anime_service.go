package service

import (
	"fmt"

	"github.com/Raditsoic/anime-go/database/repository"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnimeService struct {
	AnimeRepo     repository.AnimeRepo
	AnimeListRepo repository.AnimeListRepo
}

func NewAnimeService(animeRepo repository.AnimeRepo, animeListRepo repository.AnimeListRepo) *AnimeService {
	return &AnimeService{
		AnimeRepo:     animeRepo,
		AnimeListRepo: animeListRepo,
	}
}

func (s *AnimeService) CreateAnime(input model.NewAnime) (*model.Anime, error) {
	anime := &model.Anime{
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

func (s *AnimeService) AddAnimetoAnimeList(user_id string, entry model.NewAnimeListEntry) (*model.AnimeListEntries, error) {
	// Retrieve user AnimeList
	existingList, err := s.AnimeListRepo.GetAnimeListByUserID(user_id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If the user doesn't have a list yet, create a new one
			existingList = &model.AnimeList{
				ID:      "anime_list" + utils.GenerateUUID(),
				UserID:  user_id,
				Entries: []*model.AnimeListEntries{},
			}
		} else {
			return nil, err
		}
	}

	// Type Checking
	animeList, ok := existingList.(*model.AnimeList)
	if !ok {
		return nil, fmt.Errorf("unexpected type in AnimeListRepo: %T", existingList)
	}

	// Check if the anime is already in the list
	for _, e := range animeList.Entries {
		if e.Anime.ID == entry.ContentID {
			return nil, fmt.Errorf("Anime already in list")
		}
	}

	// Retrieve Anime Details
	entryAnime, err := s.AnimeRepo.GetByID(entry.ContentID)
	if err != nil {
		return nil, err
	}

	// Adjust NewEntry to ListEntry
	newEntry := &model.AnimeListEntries{
		ID:     "anime_list_entry" + utils.GenerateUUID(),
		Status: entry.Status,
		Anime:  entryAnime.(*model.Anime),
	}

	animeList.Entries = append(animeList.Entries, newEntry)

	err = s.AnimeListRepo.UpdateAnimeList(animeList)
	if err != nil {
		return nil, err
	}

	return newEntry, nil
}
