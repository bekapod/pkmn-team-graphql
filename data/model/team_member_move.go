package model

type TeamMemberMove struct {
	ID            string `json:"id"`
	Slot          int    `json:"slot"`
	PokemonMoveID string `json:"pokemonMoveId"`
}
