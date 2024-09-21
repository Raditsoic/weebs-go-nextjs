package graph

import "github.com/Raditsoic/anime-go/service"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{
	AnimeService *service.AnimeService
	MangaService *service.MangaService
	UserService *service.UserService
}
