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
	MovesByPokemonId dataloader.MoveListLoader
	MovesByTypeId    dataloader.MoveListLoader
	PokemonByMoveId  dataloader.PokemonListLoader
	PokemonByTypeId  dataloader.PokemonListLoader
	TypesByPokemonId dataloader.TypeListLoader
	TypeByTypeId     dataloader.TypeLoader
}

func DataLoaderMiddleware(resolver *Resolver) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &loaders{
				MovesByPokemonId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByPokemonIdDataLoader(r.Context()),
				}),
				MovesByTypeId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByTypeIdDataLoader(r.Context()),
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
				TypesByPokemonId: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypesByPokemonIdDataLoader(r.Context()),
				}),
				TypeByTypeId: *dataloader.NewTypeLoader(dataloader.TypeLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypesByTypeIdDataLoader(r.Context()),
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