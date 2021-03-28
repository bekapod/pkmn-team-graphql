package model

type Pokemon struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	PokedexId      int    `json:"pokedexId"`
	Sprite         string `json:"sprite"`
	HP             int    `json:"hp"`
	Attack         int    `json:"attack"`
	Defense        int    `json:"defense"`
	SpecialAttack  int    `json:"specialAttack"`
	SpecialDefense int    `json:"specialDefense"`
	Speed          int    `json:"speed"`
	IsBaby         bool   `json:"isBaby"`
	IsLegendary    bool   `json:"isLegendary"`
	IsMythical     bool   `json:"isMythical"`
	Description    string `json:"description"`
}

type PokemonList struct {
	Total   int        `json:"total"`
	Pokemon []*Pokemon `json:"pokemon"`
}

func NewPokemonList(pkmn []*Pokemon) PokemonList {
	return PokemonList{
		Total:   len(pkmn),
		Pokemon: pkmn,
	}
}

func NewEmptyPokemonList() PokemonList {
	return PokemonList{
		Total:   0,
		Pokemon: []*Pokemon{},
	}
}

func (p *PokemonList) AddPokemon(pkmn *Pokemon) {
	p.Total++
	p.Pokemon = append(p.Pokemon, pkmn)
}

func (Pokemon) IsEntity() {}
