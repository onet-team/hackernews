package main

import (
	"github.com/onet-team/hackernews/graph/generated"
	"github.com/onet-team/hackernews/internal/auth"

	"github.com/go-chi/chi"
	database "github.com/onet-team/hackernews/internal/pkg/db/mysql"

	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8085"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()
	server := handler.NewDefaultServer(generated.
		NewExecutableSchema(
			generated.
				Config{Resolvers: nil}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
