package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/dataloader"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
)

func (r *moveResolver) Type(ctx context.Context, obj *model.Move) (*model.Type, error) {
	return dataloader.For(ctx).TypeByTypeId.Load(obj.TypeId)
}

func (r *pokemonResolver) Types(ctx context.Context, obj *model.Pokemon) (*model.TypeList, error) {
	return dataloader.For(ctx).TypesByPokemonId.Load(obj.ID)
}

func (r *pokemonResolver) Moves(ctx context.Context, obj *model.Pokemon) (*model.MoveList, error) {
	return dataloader.For(ctx).MovesByPokemonId.Load(obj.ID)
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

func (r *queryResolver) MoveByID(ctx context.Context, id string) (*model.Move, error) {
	return r.MoveRepository.GetMoveById(ctx, id)
}

func (r *queryResolver) Moves(ctx context.Context) (*model.MoveList, error) {
	return r.MoveRepository.GetMoves(ctx)
}

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type moveResolver struct{ *Resolver }
type pokemonResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
