package graph

import (
	"bekapod/pkmn-team-graphql/data/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AbilityRepository          repository.Ability
	ItemRepository             repository.Item
	MoveRepository             repository.Move
	PokemonRepository          repository.Pokemon
	PokemonAbilityRepository   repository.PokemonAbility
	PokemonEvolutionRepository repository.PokemonEvolution
	PokemonMoveRepository      repository.PokemonMove
	PokemonTypeRepository      repository.PokemonType
	TeamRepository             repository.Team
	TypeRepository             repository.Type
}
