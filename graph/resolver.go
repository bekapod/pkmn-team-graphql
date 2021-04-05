package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"bekapod/pkmn-team-graphql/data/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AbilityRepository repository.Ability
	MoveRepository    repository.Move
	PokemonRepository repository.Pokemon
	TypeRepository    repository.Type
}
