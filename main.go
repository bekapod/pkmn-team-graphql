//go:generate go generate ./graph

package main

import (
	"bekapod/pkmn-team-graphql/config"
	"bekapod/pkmn-team-graphql/data/db"
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
	log.Logger.Info("service starting")

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Logger.Fatalf("failed to connect to client: %s", err)
	}

	resolver := &graph.Resolver{
		AbilityRepository:          repository.NewAbility(client),
		ItemRepository:             repository.NewItem(client),
		MoveRepository:             repository.NewMove(client),
		PokemonRepository:          repository.NewPokemon(client),
		PokemonAbilityRepository:   repository.NewPokemonAbility(client),
		PokemonEvolutionRepository: repository.NewPokemonEvolution(client),
		PokemonMoveRepository:      repository.NewPokemonMove(client),
		PokemonTypeRepository:      repository.NewPokemonType(client),
		TeamRepository:             repository.NewTeam(client),
		TypeRepository:             repository.NewType(client),
	}

	srv := http.Server{Addr: fmt.Sprintf(":%s", cfg.Port), Handler: router.New(resolver, cfg.Tracing)}

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop
		log.Logger.Info("shutting down HTTP server")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := client.Prisma.Disconnect(); err != nil {
			log.Logger.Errorf("Client did not disconnect correctly: %s", err)
		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Logger.Errorf("HTTP server did not shutdown correctly: %s", err)
		}
		close(httpWait)
	}()

	log.Logger.Infof("use %splayground for the GraphQL playground", cfg.AppHost)
	log.Logger.Infof("use %sgraphql for the GraphQL API", cfg.AppHost)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Logger.Fatalf("failed to start HTTP server: %s", err)
	}

	<-httpWait
}
