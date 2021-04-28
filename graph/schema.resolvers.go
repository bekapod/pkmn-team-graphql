package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
)

func (r *abilityResolver) Pokemon(ctx context.Context, obj *model.Ability) (*model.PokemonList, error) {
	pokemon, err := DataLoaderFor(ctx).PokemonByAbilityId.Load(obj.ID)

	if pokemon == nil {
		emptyPokemon := model.NewEmptyPokemonList()
		return &emptyPokemon, err
	}

	return pokemon, err
}

func (r *evolutionResolver) Pokemon(ctx context.Context, obj *model.Evolution) (*model.Pokemon, error) {
	if obj.PokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).PokemonById.Load(*obj.PokemonID)
}

func (r *evolutionResolver) KnownMove(ctx context.Context, obj *model.Evolution) (*model.Move, error) {
	if obj.KnownMoveID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).MovesById.Load(*obj.KnownMoveID)
}

func (r *evolutionResolver) KnownMoveType(ctx context.Context, obj *model.Evolution) (*model.Type, error) {
	if obj.KnownMoveTypeID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).TypesById.Load(*obj.KnownMoveTypeID)
}

func (r *evolutionResolver) PartyPokemon(ctx context.Context, obj *model.Evolution) (*model.Pokemon, error) {
	if obj.PartySpeciesPokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).PokemonById.Load(*obj.PartySpeciesPokemonID)
}

func (r *evolutionResolver) PartyPokemonType(ctx context.Context, obj *model.Evolution) (*model.Type, error) {
	if obj.PartyTypeID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).TypesById.Load(*obj.PartyTypeID)
}

func (r *evolutionResolver) TradeWithPokemon(ctx context.Context, obj *model.Evolution) (*model.Pokemon, error) {
	if obj.TradeSpeciesPokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).PokemonById.Load(*obj.TradeSpeciesPokemonID)
}

func (r *moveResolver) Pokemon(ctx context.Context, obj *model.Move) (*model.PokemonList, error) {
	pokemon, err := DataLoaderFor(ctx).PokemonByMoveId.Load(obj.ID)

	if pokemon == nil {
		emptyPokemon := model.NewEmptyPokemonList()
		return &emptyPokemon, err
	}

	return pokemon, err
}

func (r *pokemonResolver) Moves(ctx context.Context, obj *model.Pokemon) (*model.MoveList, error) {
	moves, err := DataLoaderFor(ctx).MovesByPokemonId.Load(obj.ID)

	if moves == nil {
		emptyMoves := model.NewEmptyMoveList()
		return &emptyMoves, err
	}

	return moves, err
}

func (r *pokemonResolver) EvolvesTo(ctx context.Context, obj *model.Pokemon) (*model.EvolutionList, error) {
	evolutions, err := DataLoaderFor(ctx).EvolvesToByPokemonId.Load(obj.ID)

	if evolutions == nil {
		emptyEvolutions := model.NewEmptyEvolutionList()
		return &emptyEvolutions, err
	}

	return evolutions, err
}

func (r *pokemonResolver) EvolvesFrom(ctx context.Context, obj *model.Pokemon) (*model.EvolutionList, error) {
	evolutions, err := DataLoaderFor(ctx).EvolvesFromByPokemonId.Load(obj.ID)

	if evolutions == nil {
		emptyEvolutions := model.NewEmptyEvolutionList()
		return &emptyEvolutions, err
	}

	return evolutions, err
}

func (r *queryResolver) AbilityByID(ctx context.Context, id string) (*model.Ability, error) {
	return r.AbilityRepository.GetAbilityById(ctx, id)
}

func (r *queryResolver) Abilities(ctx context.Context) (*model.AbilityList, error) {
	return r.AbilityRepository.GetAbilities(ctx)
}

func (r *queryResolver) MoveByID(ctx context.Context, id string) (*model.Move, error) {
	return r.MoveRepository.GetMoveById(ctx, id)
}

func (r *queryResolver) Moves(ctx context.Context) (*model.MoveList, error) {
	return r.MoveRepository.GetMoves(ctx)
}

func (r *queryResolver) PokemonByID(ctx context.Context, id string) (*model.Pokemon, error) {
	return r.PokemonRepository.GetPokemonById(ctx, id)
}

func (r *queryResolver) Pokemon(ctx context.Context) (*model.PokemonList, error) {
	return r.PokemonRepository.GetPokemon(ctx)
}

func (r *queryResolver) TypeByID(ctx context.Context, id string) (*model.Type, error) {
	return r.TypeRepository.GetTypeById(ctx, id)
}

func (r *queryResolver) Types(ctx context.Context) (*model.TypeList, error) {
	return r.TypeRepository.GetTypes(ctx)
}

func (r *typeResolver) Pokemon(ctx context.Context, obj *model.Type) (*model.PokemonList, error) {
	pokemon, err := DataLoaderFor(ctx).PokemonByTypeId.Load(obj.ID)

	if pokemon == nil {
		emptyPokemon := model.NewEmptyPokemonList()
		return &emptyPokemon, err
	}

	return pokemon, err
}

func (r *typeResolver) Moves(ctx context.Context, obj *model.Type) (*model.MoveList, error) {
	moves, err := DataLoaderFor(ctx).MovesByTypeId.Load(obj.ID)

	if moves == nil {
		emptyMoves := model.NewEmptyMoveList()
		return &emptyMoves, err
	}

	return moves, err
}

// Ability returns generated.AbilityResolver implementation.
func (r *Resolver) Ability() generated.AbilityResolver { return &abilityResolver{r} }

// Evolution returns generated.EvolutionResolver implementation.
func (r *Resolver) Evolution() generated.EvolutionResolver { return &evolutionResolver{r} }

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Type returns generated.TypeResolver implementation.
func (r *Resolver) Type() generated.TypeResolver { return &typeResolver{r} }

type abilityResolver struct{ *Resolver }
type evolutionResolver struct{ *Resolver }
type moveResolver struct{ *Resolver }
type pokemonResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type typeResolver struct{ *Resolver }
