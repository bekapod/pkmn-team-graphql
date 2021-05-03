package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/graph/generated"
	"context"
)

func (r *abilityResolver) Pokemon(ctx context.Context, obj *model.Ability) (*model.PokemonAbilityList, error) {
	return DataLoaderFor(ctx).PokemonAbilitiesByAbilityId.Load(obj.ID)
}

func (r *moveResolver) Type(ctx context.Context, obj *model.Move) (*model.Type, error) {
	return DataLoaderFor(ctx).TypeById.Load(obj.TypeID)
}

func (r *moveResolver) Pokemon(ctx context.Context, obj *model.Move) (*model.PokemonMoveList, error) {
	return DataLoaderFor(ctx).PokemonMovesByMoveId.Load(obj.ID)
}

func (r *pokemonResolver) Abilities(ctx context.Context, obj *model.Pokemon) (*model.PokemonAbilityList, error) {
	return DataLoaderFor(ctx).PokemonAbilitiesByPokemonId.Load(obj.ID)
}

func (r *pokemonResolver) Types(ctx context.Context, obj *model.Pokemon) (*model.PokemonTypeList, error) {
	return DataLoaderFor(ctx).PokemonTypesByPokemonId.Load(obj.ID)
}

func (r *pokemonResolver) Moves(ctx context.Context, obj *model.Pokemon) (*model.PokemonMoveList, error) {
	return DataLoaderFor(ctx).PokemonMovesByPokemonId.Load(obj.ID)
}

func (r *pokemonResolver) EvolvesTo(ctx context.Context, obj *model.Pokemon) (*model.PokemonEvolutionList, error) {
	return DataLoaderFor(ctx).PokemonEvolutionsByFromPokemonId.Load(obj.ID)
}

func (r *pokemonResolver) EvolvesFrom(ctx context.Context, obj *model.Pokemon) (*model.PokemonEvolutionList, error) {
	return DataLoaderFor(ctx).PokemonEvolutionsByToPokemonId.Load(obj.ID)
}

func (r *pokemonAbilityResolver) Ability(ctx context.Context, obj *model.PokemonAbility) (*model.Ability, error) {
	return DataLoaderFor(ctx).AbilityById.Load(obj.AbilityID)
}

func (r *pokemonAbilityResolver) Pokemon(ctx context.Context, obj *model.PokemonAbility) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).PokemonById.Load(obj.PokemonID)
}

func (r *pokemonEvolutionResolver) Pokemon(ctx context.Context, obj *model.PokemonEvolution) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).PokemonById.Load(obj.PokemonID)
}

func (r *pokemonEvolutionResolver) Item(ctx context.Context, obj *model.PokemonEvolution) (*model.Item, error) {
	if obj.ItemID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).ItemById.Load(*obj.ItemID)
}

func (r *pokemonEvolutionResolver) HeldItem(ctx context.Context, obj *model.PokemonEvolution) (*model.Item, error) {
	if obj.HeldItemID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).ItemById.Load(*obj.HeldItemID)
}

func (r *pokemonEvolutionResolver) KnownMove(ctx context.Context, obj *model.PokemonEvolution) (*model.Move, error) {
	if obj.KnownMoveID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).MoveById.Load(*obj.KnownMoveID)
}

func (r *pokemonEvolutionResolver) KnownMoveType(ctx context.Context, obj *model.PokemonEvolution) (*model.Type, error) {
	if obj.KnownMoveTypeID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).TypeById.Load(*obj.KnownMoveTypeID)
}

func (r *pokemonEvolutionResolver) PartyPokemon(ctx context.Context, obj *model.PokemonEvolution) (*model.Pokemon, error) {
	if obj.PartyPokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).PokemonById.Load(*obj.PartyPokemonID)
}

func (r *pokemonEvolutionResolver) PartyPokemonType(ctx context.Context, obj *model.PokemonEvolution) (*model.Type, error) {
	if obj.PartyTypeID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).TypeById.Load(*obj.PartyTypeID)
}

func (r *pokemonEvolutionResolver) TradeWithPokemon(ctx context.Context, obj *model.PokemonEvolution) (*model.Pokemon, error) {
	if obj.TradeWithPokemonID == nil {
		return nil, nil
	}

	return DataLoaderFor(ctx).PokemonById.Load(*obj.TradeWithPokemonID)
}

func (r *pokemonMoveResolver) Move(ctx context.Context, obj *model.PokemonMove) (*model.Move, error) {
	return DataLoaderFor(ctx).MoveById.Load(obj.MoveID)
}

func (r *pokemonMoveResolver) Pokemon(ctx context.Context, obj *model.PokemonMove) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).PokemonById.Load(obj.PokemonID)
}

func (r *pokemonTypeResolver) Type(ctx context.Context, obj *model.PokemonType) (*model.Type, error) {
	return DataLoaderFor(ctx).TypeById.Load(obj.TypeID)
}

func (r *pokemonTypeResolver) Pokemon(ctx context.Context, obj *model.PokemonType) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).PokemonById.Load(obj.PokemonID)
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

func (r *queryResolver) TeamByID(ctx context.Context, id string) (*model.Team, error) {
	return r.TeamRepository.GetTeamById(ctx, id)
}

func (r *queryResolver) Teams(ctx context.Context) (*model.TeamList, error) {
	return r.TeamRepository.GetTeams(ctx)
}

func (r *queryResolver) TypeByID(ctx context.Context, id string) (*model.Type, error) {
	return r.TypeRepository.GetTypeById(ctx, id)
}

func (r *queryResolver) Types(ctx context.Context) (*model.TypeList, error) {
	return r.TypeRepository.GetTypes(ctx)
}

func (r *teamMemberResolver) Pokemon(ctx context.Context, obj *model.TeamMember) (*model.Pokemon, error) {
	return DataLoaderFor(ctx).PokemonById.Load(obj.PokemonID)
}

func (r *teamMemberMoveResolver) Move(ctx context.Context, obj *model.TeamMemberMove) (*model.PokemonMove, error) {
	return DataLoaderFor(ctx).PokemonMoveById.Load(obj.ID)
}

func (r *typeResolver) Pokemon(ctx context.Context, obj *model.Type) (*model.PokemonTypeList, error) {
	return DataLoaderFor(ctx).PokemonTypesByTypeId.Load(obj.ID)
}

func (r *typeResolver) Moves(ctx context.Context, obj *model.Type) (*model.MoveList, error) {
	return DataLoaderFor(ctx).MovesByTypeId.Load(obj.ID)
}

func (r *typeResolver) NoDamageTo(ctx context.Context, obj *model.Type) (*model.TypeList, error) {
	return DataLoaderFor(ctx).TypesByIdWithNoDamageToRelation.Load(obj.ID)
}

func (r *typeResolver) HalfDamageTo(ctx context.Context, obj *model.Type) (*model.TypeList, error) {
	return DataLoaderFor(ctx).TypesByIdWithHalfDamageToRelation.Load(obj.ID)
}

func (r *typeResolver) DoubleDamageTo(ctx context.Context, obj *model.Type) (*model.TypeList, error) {
	return DataLoaderFor(ctx).TypesByIdWithDoubleDamageToRelation.Load(obj.ID)
}

func (r *typeResolver) NoDamageFrom(ctx context.Context, obj *model.Type) (*model.TypeList, error) {
	return DataLoaderFor(ctx).TypesByIdWithNoDamageFromRelation.Load(obj.ID)
}

func (r *typeResolver) HalfDamageFrom(ctx context.Context, obj *model.Type) (*model.TypeList, error) {
	return DataLoaderFor(ctx).TypesByIdWithHalfDamageFromRelation.Load(obj.ID)
}

func (r *typeResolver) DoubleDamageFrom(ctx context.Context, obj *model.Type) (*model.TypeList, error) {
	return DataLoaderFor(ctx).TypesByIdWithDoubleDamageFromRelation.Load(obj.ID)
}

// Ability returns generated.AbilityResolver implementation.
func (r *Resolver) Ability() generated.AbilityResolver { return &abilityResolver{r} }

// Move returns generated.MoveResolver implementation.
func (r *Resolver) Move() generated.MoveResolver { return &moveResolver{r} }

// Pokemon returns generated.PokemonResolver implementation.
func (r *Resolver) Pokemon() generated.PokemonResolver { return &pokemonResolver{r} }

// PokemonAbility returns generated.PokemonAbilityResolver implementation.
func (r *Resolver) PokemonAbility() generated.PokemonAbilityResolver {
	return &pokemonAbilityResolver{r}
}

// PokemonEvolution returns generated.PokemonEvolutionResolver implementation.
func (r *Resolver) PokemonEvolution() generated.PokemonEvolutionResolver {
	return &pokemonEvolutionResolver{r}
}

// PokemonMove returns generated.PokemonMoveResolver implementation.
func (r *Resolver) PokemonMove() generated.PokemonMoveResolver { return &pokemonMoveResolver{r} }

// PokemonType returns generated.PokemonTypeResolver implementation.
func (r *Resolver) PokemonType() generated.PokemonTypeResolver { return &pokemonTypeResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TeamMember returns generated.TeamMemberResolver implementation.
func (r *Resolver) TeamMember() generated.TeamMemberResolver { return &teamMemberResolver{r} }

// TeamMemberMove returns generated.TeamMemberMoveResolver implementation.
func (r *Resolver) TeamMemberMove() generated.TeamMemberMoveResolver {
	return &teamMemberMoveResolver{r}
}

// Type returns generated.TypeResolver implementation.
func (r *Resolver) Type() generated.TypeResolver { return &typeResolver{r} }

type abilityResolver struct{ *Resolver }
type moveResolver struct{ *Resolver }
type pokemonResolver struct{ *Resolver }
type pokemonAbilityResolver struct{ *Resolver }
type pokemonEvolutionResolver struct{ *Resolver }
type pokemonMoveResolver struct{ *Resolver }
type pokemonTypeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type teamMemberResolver struct{ *Resolver }
type teamMemberMoveResolver struct{ *Resolver }
type typeResolver struct{ *Resolver }
