package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/Raditsoic/anime-go/database"
	"github.com/Raditsoic/anime-go/graph/model"
)

var animeDatabase database.AnimeDB = *database.NewAnimeDB(database.Connect())
var mangaDatabase database.MangaDB = *database.NewMangaDB(database.Connect())

func generateUniqueID() string {
    return strconv.Itoa(rand.Int())
}

// CreateAnime is the resolver for the createAnime field.
func (r *mutationResolver) CreateAnime(ctx context.Context, input model.NewAnime) (*model.Anime, error) {
    anime := &model.Anime{
        ID:          fmt.Sprint(generateUniqueID()),
        Title:       input.Title,
        Image:       input.Image,
        Description: input.Description,
        Genre: &model.Genre{
            ID:   input.GenreID,
            Name: "Genre " + input.GenreID,
        },
    }
    animeDatabase.Create(anime)
    return anime, nil
}

func (r *mutationResolver) CreateManga(ctx context.Context, input model.NewManga) (*model.Manga, error) {
    manga := &model.Manga{
        ID:          fmt.Sprint(generateUniqueID()),
        Title:       input.Title,
        Image:       input.Image,
        Description: input.Description,
        Genre: &model.Genre{
            ID:   input.GenreID,
            Name: "Genre " + input.GenreID,
        },
    }
    mangaDatabase.Create(manga)
    return manga, nil
}

func (r *queryResolver) Animes(ctx context.Context) ([]*model.Anime, error) {
    getAnimes, err := animeDatabase.GetAll()
    if err != nil {
        log.Fatal(err)
    }

    var animes []*model.Anime
    for _, item := range getAnimes {
        anime, ok := item.(*model.Anime)
        if !ok {
            return nil, fmt.Errorf("unexpected type: %T", item)
        }
        animes = append(animes, anime)
    }

    return animes, nil
}

func (r *queryResolver) Mangas(ctx context.Context) ([]*model.Manga, error) {
    getMangas, err := mangaDatabase.GetAll()
    if err != nil {
        log.Fatal(err)
    }

    var mangas []*model.Manga
    for _, item := range getMangas {
        manga, ok := item.(*model.Manga)
        if !ok {
            return nil, fmt.Errorf("unexpected type: %T", item)
        }

        mangas = append(mangas, manga)
    }
    return mangas, nil
}

// OneAnime is the resolver for the oneAnime field.
func (r *queryResolver) OneAnime(ctx context.Context, id string) (*model.Anime, error) {
	getAnime, err := animeDatabase.GetByID(id)
    if err != nil {
        log.Fatal(err)
    }

    anime, ok := getAnime.(*model.Anime)
    if !ok {
        return nil, fmt.Errorf("unexpected type: %T", getAnime)
    }

    return anime, nil
}

// OneManga is the resolver for the oneManga field.
func (r *queryResolver) OneManga(ctx context.Context, id string) (*model.Manga, error) {
	getManga, err := mangaDatabase.GetByID(id)
    if err != nil {
        log.Fatal(err)
    }

    manga, ok := getManga.(*model.Manga)
    if !ok {
        return nil, fmt.Errorf("unexpected type: %T", getManga)
    }
    return manga, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
