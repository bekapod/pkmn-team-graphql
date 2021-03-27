package router

import (
	"bekapod/pkmn-team-graphql/graph"
	"bekapod/pkmn-team-graphql/graph/generated"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func New(resolver *graph.Resolver) *chi.Mux {
	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{"status":"OK"}`))
	})

	router.Post("/graphql", func(w http.ResponseWriter, req *http.Request) {
		gqlSrv.ServeHTTP(w, req)
	})
	router.Get("/playground", func(w http.ResponseWriter, req *http.Request) {
		playground.Handler("GraphQL", "/graphql").ServeHTTP(w, req)
	})

	return router
}
