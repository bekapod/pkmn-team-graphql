package model

type TeamMember struct {
	ID        string              `json:"id"`
	Slot      int                 `json:"slot"`
	PokemonID string              `json:"pokemonId"`
	Moves     *TeamMemberMoveList `json:"moves"`
}
