package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
)

func (r *abilityResolver) Pokemon(ctx context.Context, obj *model.Ability) (*model.PokemonWithAbilityConnection, error) {
	return DataLoaderFor(ctx).PokemonWithAbilityConnection.Load(obj.ID)
}

func (r *moveResolver) Type(ctx context.Context, obj *model.Move) (*model.Type, error) {
	return DataLoaderFor(ctx).Type.Load(obj.TypeID)
}

func (r *moveResolver) Pokemon(ctx context.Context, obj *model.Move) (*model.PokemonWithMoveConnection, error) {
	return DataLoaderFor(ctx).PokemonWithMoveConnection.Load(obj.ID)
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error) {
	return r.TeamRepository.CreateTeam(ctx, input)
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, input model.UpdateTeamInput) (*model.Team, error) {
	return r.TeamRepository.UpdateTeam(ctx, input)
}

func (r *mutationResolver) UpdateTeamMember(ctx context.Context, input model.UpdateTeamMemberInput) (*model.TeamMember, error) {
	return r.TeamRepository.UpdateTeamMember(ctx, input)
}

func (r *mutationResolver) DeleteTeam(ctx context.Context, id string) (*model.Team, error) {
	return r.TeamRepository.DeleteTeam(ctx, id)
}

func (r *mutationResolver) DeleteTeamMember(ctx context.Context, id string) (*model.TeamMember, error) {
	return r.TeamRepository.DeleteTeamMember(ctx, id)
}

func (r *mutationResolver) DeleteTeamMemberMove(ctx context.Context, id string) (*model.Move, error) {
	return r.TeamRepository.DeleteTeamMemberMove(ctx, id)
}

func (r *pokemonResolver) Abilities(ctx context.Context, obj *model.Pokemon) (*model.PokemonAbilityConnection, error) {
	return DataLoaderFor(ctx).PokemonAbilityConnection.Load(obj.ID)
}

func (r *pokemonResolver) Types(ctx context.Context, obj *model.Pokemon) (*model.PokemonTypeConnection, error) {
	return DataLoaderFor(ctx).PokemonTypeConnection.Load(obj.ID)
}

func (r *pokemonResolver) Moves(ctx context.Context, obj *model.Pokemon) (*model.PokemonMoveConnection, error) {
	return DataLoaderFor(ctx).PokemonMoveConnection.Load(obj.ID)
}

func (r *pokemonResolver) EvolvesTo(ctx context.Context, obj *model.Pokemon) (*model.PokemonEvolutionConnection, error) {
	return DataLoaderFor(ctx).PokemonEvolvesToPokemonEvolutionConnection.Load(obj.ID)
}

func (r *pokemonResolver) EvolvesFrom(ctx context.Context, obj *model.Pokemon) (*model.PokemonEvolutionConnection, error) {
	return DataLoaderFor(ctx).PokemonEvolvesFromPokemonEvolutionConnection.Load(obj.ID)
}

func (r *pokemonAbilityEdgeResolver) Node(ctx context.Context, obj *model.PokemonAbilityEdge) (*model.Ability, error) {
	return DataLoaderFor(ctx).Ability.Load(obj.NodeID)
}

func (r *pokemonEvolutionResolver) Pokemon(ctx context.Context, obj *model.PokemonEvolution) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).Pokemon.Load(obj.PokemonID)
}

func (r *pokemonEvolutionResolver) Item(ctx context.Context, obj *model.PokemonEvolution) (*model.Item, error) {
	if obj.ItemID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Item.Load(*obj.ItemID)
}

func (r *pokemonEvolutionResolver) HeldItem(ctx context.Context, obj *model.PokemonEvolution) (*model.Item, error) {
	if obj.HeldItemID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Item.Load(*obj.HeldItemID)
}

func (r *pokemonEvolutionResolver) KnownMove(ctx context.Context, obj *model.PokemonEvolution) (*model.Move, error) {
	if obj.KnownMoveID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Move.Load(*obj.KnownMoveID)
}

func (r *pokemonEvolutionResolver) KnownMoveType(ctx context.Context, obj *model.PokemonEvolution) (*model.Type, error) {
	if obj.KnownMoveTypeID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Type.Load(*obj.KnownMoveTypeID)
}

func (r *pokemonEvolutionResolver) PartyPokemon(ctx context.Context, obj *model.PokemonEvolution) (*model.Pokemon, error) {
	if obj.PartyPokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Pokemon.Load(*obj.PartyPokemonID)
}

func (r *pokemonEvolutionResolver) PartyPokemonType(ctx context.Context, obj *model.PokemonEvolution) (*model.Type, error) {
	if obj.PartyTypeID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Type.Load(*obj.PartyTypeID)
}

func (r *pokemonEvolutionResolver) TradeWithPokemon(ctx context.Context, obj *model.PokemonEvolution) (*model.Pokemon, error) {
	if obj.TradeWithPokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Pokemon.Load(*obj.TradeWithPokemonID)
}

func (r *pokemonMoveEdgeResolver) Node(ctx context.Context, obj *model.PokemonMoveEdge) (*model.Move, error) {
	return DataLoaderFor(ctx).Move.Load(obj.NodeID)
}

func (r *pokemonTypeEdgeResolver) Node(ctx context.Context, obj *model.PokemonTypeEdge) (*model.Type, error) {
	return DataLoaderFor(ctx).Type.Load(obj.NodeID)
}

func (r *pokemonWithAbilityEdgeResolver) Node(ctx context.Context, obj *model.PokemonWithAbilityEdge) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).Pokemon.Load(obj.NodeID)
}

func (r *pokemonWithMoveEdgeResolver) Node(ctx context.Context, obj *model.PokemonWithMoveEdge) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).Pokemon.Load(obj.NodeID)
}

func (r *pokemonWithTypeEdgeResolver) Node(ctx context.Context, obj *model.PokemonWithTypeEdge) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).Pokemon.Load(obj.NodeID)
}

func (r *queryResolver) AbilityByID(ctx context.Context, id string) (*model.Ability, error) {
	return r.AbilityRepository.GetAbilityById(ctx, id)
}

func (r *queryResolver) Abilities(ctx context.Context) (*model.AbilityConnection, error) {
	return r.AbilityRepository.GetAbilities(ctx)
}

func (r *queryResolver) MoveByID(ctx context.Context, id string) (*model.Move, error) {
	return r.MoveRepository.GetMoveById(ctx, id)
}

func (r *queryResolver) Moves(ctx context.Context) (*model.MoveConnection, error) {
	return r.MoveRepository.GetMoves(ctx)
}

func (r *queryResolver) PokemonByID(ctx context.Context, id string) (*model.Pokemon, error) {
	return r.PokemonRepository.GetPokemonById(ctx, id)
}

func (r *queryResolver) Pokemon(ctx context.Context) (*model.PokemonConnection, error) {
	return r.PokemonRepository.GetPokemon(ctx)
}

func (r *queryResolver) TeamByID(ctx context.Context, id string) (*model.Team, error) {
	return r.TeamRepository.GetTeamById(ctx, id)
}

func (r *queryResolver) Teams(ctx context.Context) (*model.TeamConnection, error) {
	return r.TeamRepository.GetTeams(ctx)
}

func (r *queryResolver) TypeByID(ctx context.Context, id string) (*model.Type, error) {
	return r.TypeRepository.GetTypeById(ctx, id)
}

func (r *queryResolver) Types(ctx context.Context) (*model.TypeConnection, error) {
	return r.TypeRepository.GetTypes(ctx)
}

func (r *teamMemberResolver) Pokemon(ctx context.Context, obj *model.TeamMember) (*model.Pokemon, error) {
	if obj.PokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).Pokemon.Load(*obj.PokemonID)
}

func (r *teamMemberMoveEdgeResolver) Node(ctx context.Context, obj *model.TeamMemberMoveEdge) (*model.Move, error) {
	return DataLoaderFor(ctx).Move.Load(obj.NodeID)
}

func (r *typeResolver) Pokemon(ctx context.Context, obj *model.Type) (*model.PokemonWithTypeConnection, error) {
	return DataLoaderFor(ctx).PokemonWithTypeConnection.Load(obj.ID)
}

func (r *typeResolver) Moves(ctx context.Context, obj *model.Type) (*model.MoveConnection, error) {
	return DataLoaderFor(ctx).MovesWithType.Load(obj.ID)
}

func (r *typeResolver) NoDamageTo(ctx context.Context, obj *model.Type) (*model.TypeConnection, error) {
	return DataLoaderFor(ctx).TypeNoDamageToTypesConnection.Load(obj.ID)
}

func (r *typeResolver) HalfDamageTo(ctx context.Context, obj *model.Type) (*model.TypeConnection, error) {
	return DataLoaderFor(ctx).TypeHalfDamageToTypesConnection.Load(obj.ID)
}

func (r *typeResolver) DoubleDamageTo(ctx context.Context, obj *model.Type) (*model.TypeConnection, error) {
	return DataLoaderFor(ctx).TypeDoubleDamageToTypesConnection.Load(obj.ID)
}

func (r *typeResolver) NoDamageFrom(ctx context.Context, obj *model.Type) (*model.TypeConnection, error) {
	return DataLoaderFor(ctx).TypeNoDamageFromTypesConnection.Load(obj.ID)
}

func (r *typeResolver) HalfDamageFrom(ctx context.Context, obj *model.Type) (*model.TypeConnection, error) {
	return DataLoaderFor(ctx).TypeHalfDamageFromTypesConnection.Load(obj.ID)
}

func (r *typeResolver) DoubleDamageFrom(ctx context.Context, obj *model.Type) (*model.TypeConnection, error) {
	return DataLoaderFor(ctx).TypeDoubleDamageFromTypesConnection.Load(obj.ID)
}

// Ability returns generated.AbilityResolver implementation.
func (r *Resolver) Ability() generated.AbilityResolver { return &abilityResolver{r} }

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// PokemonAbilityEdge returns generated.PokemonAbilityEdgeResolver implementation.
func (r *Resolver) PokemonAbilityEdge() generated.PokemonAbilityEdgeResolver {
	return &pokemonAbilityEdgeResolver{r}
}

// PokemonEvolution returns generated.PokemonEvolutionResolver implementation.
func (r *Resolver) PokemonEvolution() generated.PokemonEvolutionResolver {
	return &pokemonEvolutionResolver{r}
}

// PokemonMoveEdge returns generated.PokemonMoveEdgeResolver implementation.
func (r *Resolver) PokemonMoveEdge() generated.PokemonMoveEdgeResolver {
	return &pokemonMoveEdgeResolver{r}
}

// PokemonTypeEdge returns generated.PokemonTypeEdgeResolver implementation.
func (r *Resolver) PokemonTypeEdge() generated.PokemonTypeEdgeResolver {
	return &pokemonTypeEdgeResolver{r}
}

// PokemonWithAbilityEdge returns generated.PokemonWithAbilityEdgeResolver implementation.
func (r *Resolver) PokemonWithAbilityEdge() generated.PokemonWithAbilityEdgeResolver {
	return &pokemonWithAbilityEdgeResolver{r}
}

// PokemonWithMoveEdge returns generated.PokemonWithMoveEdgeResolver implementation.
func (r *Resolver) PokemonWithMoveEdge() generated.PokemonWithMoveEdgeResolver {
	return &pokemonWithMoveEdgeResolver{r}
}

// PokemonWithTypeEdge returns generated.PokemonWithTypeEdgeResolver implementation.
func (r *Resolver) PokemonWithTypeEdge() generated.PokemonWithTypeEdgeResolver {
	return &pokemonWithTypeEdgeResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TeamMember returns generated.TeamMemberResolver implementation.
func (r *Resolver) TeamMember() generated.TeamMemberResolver { return &teamMemberResolver{r} }

// TeamMemberMoveEdge returns generated.TeamMemberMoveEdgeResolver implementation.
func (r *Resolver) TeamMemberMoveEdge() generated.TeamMemberMoveEdgeResolver {
	return &teamMemberMoveEdgeResolver{r}
}

// Type returns generated.TypeResolver implementation.
func (r *Resolver) Type() generated.TypeResolver { return &typeResolver{r} }

type abilityResolver struct{ *Resolver }
type moveResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type pokemonResolver struct{ *Resolver }
type pokemonAbilityEdgeResolver struct{ *Resolver }
type pokemonEvolutionResolver struct{ *Resolver }
type pokemonMoveEdgeResolver struct{ *Resolver }
type pokemonTypeEdgeResolver struct{ *Resolver }
type pokemonWithAbilityEdgeResolver struct{ *Resolver }
type pokemonWithMoveEdgeResolver struct{ *Resolver }
type pokemonWithTypeEdgeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type teamMemberResolver struct{ *Resolver }
type teamMemberMoveEdgeResolver struct{ *Resolver }
type typeResolver struct{ *Resolver }
