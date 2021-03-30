package router

import (
	"bekapod/pkmn-team-graphql/graph"
	"bekapod/pkmn-team-graphql/graph/generated"
	"bekapod/pkmn-team-graphql/log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func New(resolver *graph.Resolver, tracing bool) *chi.Mux {
	c := generated.Config{Resolvers: resolver}
	countComplexity := func(childComplexity int) int {
		return childComplexity * childComplexity
	}
	c.Complexity.Pokemon.Moves = countComplexity
	c.Complexity.Pokemon.Types = countComplexity
	c.Complexity.Move.Type = countComplexity
	c.Complexity.Move.Pokemon = countComplexity
	c.Complexity.Type.Moves = countComplexity
	c.Complexity.Type.Pokemon = countComplexity
	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	gqlSrv.Use(extension.FixedComplexityLimit(500))

	if tracing {
		log.Logger.Info("using apollo tracing")
		gqlSrv.Use(apollotracing.Tracer{})
	}

	router := chi.NewRouter()
	router.Use(graph.DataLoaderMiddleware(resolver))

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
