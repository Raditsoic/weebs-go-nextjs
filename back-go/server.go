package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Raditsoic/anime-go/database"
	"github.com/Raditsoic/anime-go/graph"
	"github.com/Raditsoic/anime-go/service"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	animeRepo := database.NewAnimeRepo(database.Connect())
	mangaRepo := database.NewMangaRepo(database.Connect())

	animeService := service.NewAnimeService(*animeRepo)
	mangaService := service.NewMangaService(*mangaRepo)

	resolver := &graph.Resolver{
		AnimeService: animeService,
		MangaService: mangaService,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	corsMiddleware := cors.Default()

	router := http.NewServeMux()

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	handlerWithCors := corsMiddleware.Handler(router)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCors))
}
