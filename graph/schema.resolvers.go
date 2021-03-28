package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/dataloader"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
	"fmt"
)

func (r *moveResolver) Type(ctx context.Context, obj *model.Move) (*model.Type, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *pokemonResolver) Types(ctx context.Context, obj *model.Pokemon) (*model.TypeList, error) {
	return dataloader.For(ctx).TypesByPokemonId.Load(obj.ID)
}

func (r *pokemonResolver) Moves(ctx context.Context, obj *model.Pokemon) (*model.MoveList, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPokemon(ctx context.Context) (*model.PokemonList, error) {
	return r.PokemonRepository.GetAllPokemon(ctx)
}

func (r *queryResolver) AllTypes(ctx context.Context) (*model.TypeList, error) {
	return r.TypeRepository.GetAllTypes(ctx)
}

func (r *queryResolver) AllMoves(ctx context.Context) (*model.MoveList, error) {
	panic(fmt.Errorf("not implemented"))
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
