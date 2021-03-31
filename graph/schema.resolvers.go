package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
	"fmt"
)

func (r *abilityResolver) Pokemon(ctx context.Context, obj *model.Ability) (*model.PokemonList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *moveResolver) Type(ctx context.Context, obj *model.Move) (*model.Type, error) {
	return DataLoaderFor(ctx).TypeByTypeId.Load(obj.TypeID)
}

func (r *moveResolver) Pokemon(ctx context.Context, obj *model.Move) (*model.PokemonList, error) {
	pokemon, err := DataLoaderFor(ctx).PokemonByMoveId.Load(obj.ID)

	if pokemon == nil {
		emptyPokemon := model.NewEmptyPokemonList()
		return &emptyPokemon, err
	}

	return pokemon, err
}

func (r *pokemonResolver) Abilities(ctx context.Context, obj *model.Pokemon) (*model.AbilityList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pokemonResolver) Types(ctx context.Context, obj *model.Pokemon) (*model.TypeList, error) {
	types, err := DataLoaderFor(ctx).TypesByPokemonId.Load(obj.ID)

	if types == nil {
		emptyTypes := model.NewEmptyTypeList()
		return &emptyTypes, err
	}

	return types, err
}

func (r *pokemonResolver) Moves(ctx context.Context, obj *model.Pokemon) (*model.MoveList, error) {
	moves, err := DataLoaderFor(ctx).MovesByPokemonId.Load(obj.ID)

	if moves == nil {
		emptyMoves := model.NewEmptyMoveList()
		return &emptyMoves, err
	}

	return moves, err
}

func (r *queryResolver) AbilityByID(ctx context.Context, id string) (*model.Ability, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Abilities(ctx context.Context) (*model.AbilityList, error) {
	panic(fmt.Errorf("not implemented"))
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

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Type returns generated.TypeResolver implementation.
func (r *Resolver) Type() generated.TypeResolver { return &typeResolver{r} }

type abilityResolver struct{ *Resolver }
type moveResolver struct{ *Resolver }
type pokemonResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type typeResolver struct{ *Resolver }
