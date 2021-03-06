package router

import (
	"bekapod/pkmn-team-graphql/graph"
	"bekapod/pkmn-team-graphql/graph/generated"
	"bekapod/pkmn-team-graphql/log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func New(resolver *graph.Resolver, tracing bool) *chi.Mux {
	c := generated.Config{Resolvers: resolver}
	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

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
