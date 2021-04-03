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
	AbilitiesByPokemonId dataloader.AbilityListLoader
	EggGroupsByPokemonId dataloader.EggGroupListLoader
	MovesByPokemonId     dataloader.MoveListLoader
	MovesByTypeId        dataloader.MoveListLoader
	PokemonByAbilityId   dataloader.PokemonListLoader
	PokemonByMoveId      dataloader.PokemonListLoader
	PokemonByTypeId      dataloader.PokemonListLoader
}

func DataLoaderMiddleware(resolver *Resolver) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &loaders{
				AbilitiesByPokemonId: *dataloader.NewAbilityListLoader(dataloader.AbilityListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.AbilityRepository.AbilitiesByPokemonIdDataLoader(r.Context()),
				}),
				EggGroupsByPokemonId: *dataloader.NewEggGroupListLoader(dataloader.EggGroupListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.EggGroupRepository.EggGroupsByPokemonIdDataLoader(r.Context()),
				}),
				MovesByPokemonId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByPokemonIdDataLoader(r.Context()),
				}),
				MovesByTypeId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByTypeIdDataLoader(r.Context()),
				}),
				PokemonByAbilityId: *dataloader.NewPokemonListLoader(dataloader.PokemonListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByAbilityIdDataLoader(r.Context()),
				}),
				PokemonByMoveId: *dataloader.NewPokemonListLoader(dataloader.PokemonListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByMoveIdDataLoader(r.Context()),
				}),
				PokemonByTypeId: *dataloader.NewPokemonListLoader(dataloader.PokemonListLoaderConfig{
					MaxBatch: 1000,
					Wait:     10 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByTypeIdDataLoader(r.Context()),
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
