package service

import (
	"fmt"

	"github.com/Raditsoic/anime-go/database/repository"
	"github.com/Raditsoic/anime-go/graph/model"
	"github.com/Raditsoic/anime-go/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type MangaService struct {
	MangaRepo     repository.MangaRepo
	MangaListRepo repository.MangaListRepo
}

func NewMangaService(mangaRepo repository.MangaRepo, mangaListRepo repository.MangaListRepo) *MangaService {
	return &MangaService{
		MangaRepo:     mangaRepo,
		MangaListRepo: mangaListRepo}
}

func (s *MangaService) CreateManga(input model.NewManga) (*model.Manga, error) {
	manga := &model.Manga{
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

func (s *MangaService) AddMangatoMangaList(user_id string, entry model.NewMangaListEntry) (*model.MangaListEntries, error) {
	// Retrieve user MangaList
	existingList, err := s.MangaListRepo.GetMangaListByUserID(user_id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If the user doesn't have a list yet, create a new one
			existingList = &model.MangaList{
				ID: "manga_list" + utils.GenerateUUID(),
				UserID: user_id,
				Entries: []*model.MangaListEntries{},
			}
		} else {
			return nil, err
		}
	}

	// Type Checking
	mangaList, ok := existingList.(*model.MangaList)
	if !ok {
		return nil, fmt.Errorf("unexpected type in MangaListRepo: %T", existingList)
	}

	// Check if the manga is already in the list
	for _, e := range mangaList.Entries {
		if e.Manga.ID == entry.ContentID {
			return nil, fmt.Errorf("Manga already in list")
		}
	}

	// Retrieve Manga Details
	entryManga, err := s.MangaRepo.GetByID(entry.ContentID)
	if err != nil {
		return nil, err
	}

	// Adjust NewEntry to ListEntry
	newEntry := &model.MangaListEntries{
		ID:       "manga_list_entry" + utils.GenerateUUID(),
		Status:    entry.Status,
		Manga:    entryManga.(*model.Manga),
	}

	mangaList.Entries = append(mangaList.Entries, newEntry)
	
	err = s.MangaListRepo.UpdateMangaList(mangaList)
	if err != nil {
		return nil, err
	}

	return newEntry, nil
}
