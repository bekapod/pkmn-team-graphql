package model

import "bekapod/pkmn-team-graphql/data/db"

type Pokemon struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Slug             string              `json:"slug"`
	PokedexID        int                 `json:"pokedexId"`
	Sprite           *string             `json:"sprite"`
	Hp               int                 `json:"hp"`
	Attack           int                 `json:"attack"`
	Defense          int                 `json:"defense"`
	SpecialAttack    int                 `json:"specialAttack"`
	SpecialDefense   int                 `json:"specialDefense"`
	Speed            int                 `json:"speed"`
	IsBaby           bool                `json:"isBaby"`
	IsLegendary      bool                `json:"isLegendary"`
	IsMythical       bool                `json:"isMythical"`
	Description      *string             `json:"description"`
	Color            Color               `json:"color"`
	Shape            Shape               `json:"shape"`
	Habitat          *Habitat            `json:"habitat"`
	Height           int                 `json:"height"`
	Weight           int                 `json:"weight"`
	IsDefaultVariant bool                `json:"isDefaultVariant"`
	Genus            string              `json:"genus"`
	EggGroups        *EggGroupConnection `json:"eggGroups"`
}

func (Pokemon) IsNode() {}

func NewPokemonFromDb(dbPokemon db.PokemonModel) Pokemon {
	p := Pokemon{
		ID:               dbPokemon.ID,
		Slug:             dbPokemon.Slug,
		Name:             dbPokemon.Name,
		PokedexID:        dbPokemon.PokedexID,
		Hp:               dbPokemon.Hp,
		Attack:           dbPokemon.Attack,
		Defense:          dbPokemon.Defense,
		SpecialAttack:    dbPokemon.SpecialAttack,
		SpecialDefense:   dbPokemon.SpecialDefense,
		Speed:            dbPokemon.Speed,
		IsBaby:           dbPokemon.IsBaby,
		IsLegendary:      dbPokemon.IsLegendary,
		IsMythical:       dbPokemon.IsMythical,
		Color:            Color(dbPokemon.Color),
		Shape:            Shape(dbPokemon.Shape),
		Height:           dbPokemon.Height,
		Weight:           dbPokemon.Weight,
		IsDefaultVariant: dbPokemon.IsDefaultVariant,
		Genus:            dbPokemon.Genus,
	}

	if sprite, ok := dbPokemon.Sprite(); ok {
		p.Sprite = &sprite
	} else {
		p.Sprite = nil
	}

	if habitat, ok := dbPokemon.Habitat(); ok {
		h := Habitat(habitat)
		p.Habitat = &h
	} else {
		p.Habitat = nil
	}

	if description, ok := dbPokemon.Description(); ok {
		p.Description = &description
	} else {
		p.Description = nil
	}

	eggGroups := NewEmptyEggGroupConnection()
	for _, e := range dbPokemon.EggGroups() {
		eggGroup := NewEggGroupEdgeFromDb(e)
		eggGroups.AddEdge(&eggGroup)
	}

	p.EggGroups = &eggGroups

	return p
}

func NewPokemonEdgeFromDb(dbPokemon db.PokemonModel) PokemonEdge {
	node := NewPokemonFromDb(dbPokemon)
	return PokemonEdge{
		Cursor: dbPokemon.ID,
		Node:   &node,
	}
}

func NewEmptyPokemonConnection() PokemonConnection {
	return PokemonConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonEdge{},
	}
}

func (c *PokemonConnection) AddEdge(e *PokemonEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
