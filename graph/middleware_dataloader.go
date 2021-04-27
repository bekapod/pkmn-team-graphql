package graph

import (
	"bekapod/pkmn-team-graphql/dataloader"
	"context"
	"net/http"
	"time"
)

type key string

const loadersKey key = "dataloaders"

type loaders struct {
	EvolutionsByPokemonId dataloader.EvolutionListLoader
	MovesById             dataloader.MoveLoader
	MovesByPokemonId      dataloader.MoveListLoader
	MovesByTypeId         dataloader.MoveListLoader
	PokemonById           dataloader.PokemonLoader
	PokemonByAbilityId    dataloader.PokemonListLoader
	PokemonByMoveId       dataloader.PokemonListLoader
	PokemonByTypeId       dataloader.PokemonListLoader
	TypesById             dataloader.TypeLoader
}

func DataLoaderMiddleware(resolver *Resolver) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &loaders{
				EvolutionsByPokemonId: *dataloader.NewEvolutionListLoader(dataloader.EvolutionListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.EvolutionRepository.EvolutionsByPokemonIdDataLoader(r.Context()),
				}),
				MovesById: *dataloader.NewMoveLoader(dataloader.MoveLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByIdDataLoader(r.Context()),
				}),
				MovesByPokemonId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByPokemonIdDataLoader(r.Context()),
				}),
				MovesByTypeId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByTypeIdDataLoader(r.Context()),
				}),
				PokemonById: *dataloader.NewPokemonLoader(dataloader.PokemonLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByIdDataLoader(r.Context()),
				}),
				PokemonByAbilityId: *dataloader.NewPokemonListLoader(dataloader.PokemonListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByAbilityIdDataLoader(r.Context()),
				}),
				PokemonByMoveId: *dataloader.NewPokemonListLoader(dataloader.PokemonListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByMoveIdDataLoader(r.Context()),
				}),
				PokemonByTypeId: *dataloader.NewPokemonListLoader(dataloader.PokemonListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByTypeIdDataLoader(r.Context()),
				}),
				TypesById: *dataloader.NewTypeLoader(dataloader.TypeLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypesByIdDataLoader(r.Context()),
				}),
			})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func DataLoaderFor(ctx context.Context) *loaders {
	return ctx.Value(loadersKey).(*loaders)
}
