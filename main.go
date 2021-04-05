//go:generate go generate ./graph

package main

import (
	"bekapod/pkmn-team-graphql/config"
	"bekapod/pkmn-team-graphql/data"
	"bekapod/pkmn-team-graphql/data/repository"
	"bekapod/pkmn-team-graphql/graph"
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	httpWait := make(chan struct{})
	cfg := config.New()
	log.Logger.WithField("config", cfg).Info("service starting")

	db := data.NewDB(cfg.DatabaseUrl)
	data.MigrateDatabase(db, cfg)

	resolver := &graph.Resolver{
		AbilityRepository: repository.NewAbility(db),
		MoveRepository:    repository.NewMove(db),
		PokemonRepository: repository.NewPokemon(db),
		TypeRepository:    repository.NewType(db),
	}

	srv := http.Server{Addr: fmt.Sprintf(":%s", cfg.Port), Handler: router.New(resolver, cfg.Tracing)}

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop
		log.Logger.Info("shutting down HTTP server")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Logger.Errorf("HTTP server did not shutdown correctly: %s", err)
		}
		close(httpWait)
	}()

	log.Logger.Infof("use https://%s/playground for the GraphQL playground", cfg.AppHostWithPort())
	log.Logger.Infof("use https://%s/graphql for the GraphQL API", cfg.AppHostWithPort())

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Logger.Fatalf("failed to start HTTP server: %s", err)
	}

	<-httpWait
}
