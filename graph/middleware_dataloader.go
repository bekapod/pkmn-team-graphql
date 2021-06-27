package graph

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/dataloader"
	"context"
	"net/http"
	"time"
)

type key string

const loadersKey key = "dataloaders"

type loaders struct {
	Ability                                      dataloader.AbilityLoader
	Item                                         dataloader.ItemLoader
	MovesWithType                                dataloader.MoveConnectionLoader
	Move                                         dataloader.MoveLoader
	PokemonAbilityConnection                     dataloader.PokemonAbilityConnectionLoader
	PokemonEvolvesToPokemonEvolutionConnection   dataloader.PokemonEvolutionConnectionLoader
	PokemonEvolvesFromPokemonEvolutionConnection dataloader.PokemonEvolutionConnectionLoader
	Pokemon                                      dataloader.PokemonLoader
	PokemonMoveConnection                        dataloader.PokemonMoveConnectionLoader
	PokemonTypeConnection                        dataloader.PokemonTypeConnectionLoader
	PokemonWithAbilityConnection                 dataloader.PokemonWithAbilityConnectionLoader
	PokemonWithMoveConnection                    dataloader.PokemonWithMoveConnectionLoader
	PokemonWithTypeConnection                    dataloader.PokemonWithTypeConnectionLoader
	TypeNoDamageToTypesConnection                dataloader.TypeConnectionLoader
	TypeHalfDamageToTypesConnection              dataloader.TypeConnectionLoader
	TypeDoubleDamageToTypesConnection            dataloader.TypeConnectionLoader
	TypeNoDamageFromTypesConnection              dataloader.TypeConnectionLoader
	TypeHalfDamageFromTypesConnection            dataloader.TypeConnectionLoader
	TypeDoubleDamageFromTypesConnection          dataloader.TypeConnectionLoader
	Type                                         dataloader.TypeLoader
}

func DataLoaderMiddleware(resolver *Resolver) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &loaders{
				Ability: *dataloader.NewAbilityLoader(dataloader.AbilityLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.AbilityRepository.AbilityByIdDataLoader(r.Context()),
				}),
				Item: *dataloader.NewItemLoader(dataloader.ItemLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.ItemRepository.ItemByIdDataLoader(r.Context()),
				}),
				MovesWithType: *dataloader.NewMoveConnectionLoader(dataloader.MoveConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByTypeIdDataLoader(r.Context()),
				}),
				Move: *dataloader.NewMoveLoader(dataloader.MoveLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MoveByIdDataLoader(r.Context()),
				}),
				PokemonAbilityConnection: *dataloader.NewPokemonAbilityConnectionLoader(dataloader.PokemonAbilityConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonAbilityRepository.PokemonAbilityByPokemonIdDataLoader(r.Context()),
				}),
				PokemonEvolvesToPokemonEvolutionConnection: *dataloader.NewPokemonEvolutionConnectionLoader(dataloader.PokemonEvolutionConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonEvolutionRepository.PokemonEvolutionByFromPokemonIdDataLoader(r.Context()),
				}),
				PokemonEvolvesFromPokemonEvolutionConnection: *dataloader.NewPokemonEvolutionConnectionLoader(dataloader.PokemonEvolutionConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonEvolutionRepository.PokemonEvolutionByToPokemonIdDataLoader(r.Context()),
				}),
				Pokemon: *dataloader.NewPokemonLoader(dataloader.PokemonLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByIdDataLoader(r.Context()),
				}),
				PokemonMoveConnection: *dataloader.NewPokemonMoveConnectionLoader(dataloader.PokemonMoveConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonMoveRepository.PokemonMoveByPokemonIdDataLoader(r.Context()),
				}),
				PokemonTypeConnection: *dataloader.NewPokemonTypeConnectionLoader(dataloader.PokemonTypeConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonTypeRepository.PokemonTypeByPokemonIdDataLoader(r.Context()),
				}),
				PokemonWithAbilityConnection: *dataloader.NewPokemonWithAbilityConnectionLoader(dataloader.PokemonWithAbilityConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonAbilityRepository.PokemonAbilityByAbilityIdDataLoader(r.Context()),
				}),
				PokemonWithMoveConnection: *dataloader.NewPokemonWithMoveConnectionLoader(dataloader.PokemonWithMoveConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonMoveRepository.PokemonMoveByMoveIdDataLoader(r.Context()),
				}),
				PokemonWithTypeConnection: *dataloader.NewPokemonWithTypeConnectionLoader(dataloader.PokemonWithTypeConnectionLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonTypeRepository.PokemonTypeByTypeIdDataLoader(r.Context()),
				}),
				TypeNoDamageToTypesConnection: *dataloader.NewTypeConnectionLoader(dataloader.TypeConnectionLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationNODAMAGETO),
				}),
				TypeHalfDamageToTypesConnection: *dataloader.NewTypeConnectionLoader(dataloader.TypeConnectionLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationHALFDAMAGETO),
				}),
				TypeDoubleDamageToTypesConnection: *dataloader.NewTypeConnectionLoader(dataloader.TypeConnectionLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationDOUBLEDAMAGETO),
				}),
				TypeNoDamageFromTypesConnection: *dataloader.NewTypeConnectionLoader(dataloader.TypeConnectionLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationNODAMAGEFROM),
				}),
				TypeHalfDamageFromTypesConnection: *dataloader.NewTypeConnectionLoader(dataloader.TypeConnectionLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationHALFDAMAGEFROM),
				}),
				TypeDoubleDamageFromTypesConnection: *dataloader.NewTypeConnectionLoader(dataloader.TypeConnectionLoaderConfig{
					MaxBatch: 1000,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationDOUBLEDAMAGEFROM),
				}),
				Type: *dataloader.NewTypeLoader(dataloader.TypeLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdDataLoader(r.Context()),
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
