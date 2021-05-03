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
	AbilityById                           dataloader.AbilityLoader
	ItemById                              dataloader.ItemLoader
	MoveById                              dataloader.MoveLoader
	MovesByTypeId                         dataloader.MoveListLoader
	PokemonById                           dataloader.PokemonLoader
	PokemonAbilitiesByAbilityId           dataloader.PokemonAbilityListLoader
	PokemonAbilitiesByPokemonId           dataloader.PokemonAbilityListLoader
	PokemonEvolutionsByFromPokemonId      dataloader.PokemonEvolutionListLoader
	PokemonEvolutionsByToPokemonId        dataloader.PokemonEvolutionListLoader
	PokemonMoveById                       dataloader.PokemonMoveLoader
	PokemonMovesByMoveId                  dataloader.PokemonMoveListLoader
	PokemonMovesByPokemonId               dataloader.PokemonMoveListLoader
	PokemonTypesByPokemonId               dataloader.PokemonTypeListLoader
	PokemonTypesByTypeId                  dataloader.PokemonTypeListLoader
	TypeById                              dataloader.TypeLoader
	TypesByIdWithNoDamageToRelation       dataloader.TypeListLoader
	TypesByIdWithHalfDamageToRelation     dataloader.TypeListLoader
	TypesByIdWithDoubleDamageToRelation   dataloader.TypeListLoader
	TypesByIdWithNoDamageFromRelation     dataloader.TypeListLoader
	TypesByIdWithHalfDamageFromRelation   dataloader.TypeListLoader
	TypesByIdWithDoubleDamageFromRelation dataloader.TypeListLoader
}

func DataLoaderMiddleware(resolver *Resolver) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &loaders{
				AbilityById: *dataloader.NewAbilityLoader(dataloader.AbilityLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.AbilityRepository.AbilityByIdDataLoader(r.Context()),
				}),
				ItemById: *dataloader.NewItemLoader(dataloader.ItemLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.ItemRepository.ItemByIdDataLoader(r.Context()),
				}),
				MoveById: *dataloader.NewMoveLoader(dataloader.MoveLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MoveByIdDataLoader(r.Context()),
				}),
				MovesByTypeId: *dataloader.NewMoveListLoader(dataloader.MoveListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.MoveRepository.MovesByTypeIdDataLoader(r.Context()),
				}),
				PokemonById: *dataloader.NewPokemonLoader(dataloader.PokemonLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonRepository.PokemonByIdDataLoader(r.Context()),
				}),
				PokemonAbilitiesByAbilityId: *dataloader.NewPokemonAbilityListLoader(dataloader.PokemonAbilityListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonAbilityRepository.PokemonAbilityByAbilityIdDataLoader(r.Context()),
				}),
				PokemonAbilitiesByPokemonId: *dataloader.NewPokemonAbilityListLoader(dataloader.PokemonAbilityListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonAbilityRepository.PokemonAbilityByPokemonIdDataLoader(r.Context()),
				}),
				PokemonEvolutionsByFromPokemonId: *dataloader.NewPokemonEvolutionListLoader(dataloader.PokemonEvolutionListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonEvolutionRepository.PokemonEvolutionByFromPokemonIdDataLoader(r.Context()),
				}),
				PokemonEvolutionsByToPokemonId: *dataloader.NewPokemonEvolutionListLoader(dataloader.PokemonEvolutionListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonEvolutionRepository.PokemonEvolutionByToPokemonIdDataLoader(r.Context()),
				}),
				PokemonMoveById: *dataloader.NewPokemonMoveLoader(dataloader.PokemonMoveLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonMoveRepository.PokemonMoveByIdDataLoader(r.Context()),
				}),
				PokemonMovesByMoveId: *dataloader.NewPokemonMoveListLoader(dataloader.PokemonMoveListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonMoveRepository.PokemonMoveByMoveIdDataLoader(r.Context()),
				}),
				PokemonMovesByPokemonId: *dataloader.NewPokemonMoveListLoader(dataloader.PokemonMoveListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonMoveRepository.PokemonMoveByPokemonIdDataLoader(r.Context()),
				}),
				PokemonTypesByPokemonId: *dataloader.NewPokemonTypeListLoader(dataloader.PokemonTypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonTypeRepository.PokemonTypeByPokemonIdDataLoader(r.Context()),
				}),
				PokemonTypesByTypeId: *dataloader.NewPokemonTypeListLoader(dataloader.PokemonTypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.PokemonTypeRepository.PokemonTypeByTypeIdDataLoader(r.Context()),
				}),
				TypeById: *dataloader.NewTypeLoader(dataloader.TypeLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdDataLoader(r.Context()),
				}),
				TypesByIdWithNoDamageToRelation: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationNODAMAGETO),
				}),
				TypesByIdWithHalfDamageToRelation: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationHALFDAMAGETO),
				}),
				TypesByIdWithDoubleDamageToRelation: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationDOUBLEDAMAGETO),
				}),
				TypesByIdWithNoDamageFromRelation: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationNODAMAGEFROM),
				}),
				TypesByIdWithHalfDamageFromRelation: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationHALFDAMAGEFROM),
				}),
				TypesByIdWithDoubleDamageFromRelation: *dataloader.NewTypeListLoader(dataloader.TypeListLoaderConfig{
					MaxBatch: 0,
					Wait:     1 * time.Millisecond,
					Fetch:    resolver.TypeRepository.TypeByIdWithDamageRelationDataLoader(r.Context(), db.DamageRelationDOUBLEDAMAGEFROM),
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
