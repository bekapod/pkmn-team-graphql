package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
	"fmt"
)

func (r *pokemonResolver) Types(ctx context.Context, obj *model.Pokemon) ([]*model.Type, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPokemon(ctx context.Context) (*model.PokemonList, error) {
	return r.PokemonRepository.GetAllPokemon(ctx)
}

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type pokemonResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
