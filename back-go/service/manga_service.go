package service

import (
	"fmt"

	"github.com/Raditsoic/anime-go/database"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
)

type MangaService struct {
	MangaRepo database.MangaRepo
}

func NewMangaService(repo database.MangaRepo) *MangaService {
	return &MangaService{MangaRepo: repo}
}

func (s *MangaService) CreateManga(input model.NewManga) (*model.Manga, error) {
	manga := &model.Manga{
		ID:          fmt.Sprint(utils.GenerateUniqueID()),
		Title:       input.Title,
		Image:       input.Image,
		Description: input.Description,
		Genre: &model.Genre{
			ID:   input.GenreID,
			Name: "Genre " + input.GenreID,
		},
	}
	if err := s.MangaRepo.Create(manga); err != nil {
		return nil, err
	}
	return manga, nil
}

func (s *MangaService) GetMangas() ([]*model.Manga, error) {
	mangas, err := s.MangaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	mangaList := make([]*model.Manga, len(mangas))
	for i, item := range mangas {
		manga, ok := item.(*model.Manga)
		if !ok {
			return nil, fmt.Errorf("unexpected type in AnimeRepo: %T", item)
		}
		mangaList[i] = manga
	}
	return mangaList, nil
}

func (s *MangaService) GetMangaByID(id string) (*model.Manga, error) {
    mangaInterface, err := s.MangaRepo.GetByID(id)
    if err != nil {
        return nil, err
    }

    manga, ok := mangaInterface.(*model.Manga)
    if !ok {
        return nil, fmt.Errorf("unexpected type in MangaRepo: %T", mangaInterface)
    }
    return manga, nil
}
