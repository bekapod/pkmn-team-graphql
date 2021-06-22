package model

import "bekapod/pkmn-team-graphql/data/db"

type PokemonEvolution struct {
	ID                    string           `json:"id"`
	FromPokemonID         string           `json:"fromPokemonId"`
	ToPokemonID           string           `json:"toPokemonId"`
	PokemonID             string           `json:"pokemonId"`
	Trigger               EvolutionTrigger `json:"trigger"`
	ItemID                *string          `json:"itemId"`
	Gender                Gender           `json:"gender"`
	HeldItemID            *string          `json:"heldItemId"`
	KnownMoveID           *string          `json:"knownMoveId"`
	KnownMoveTypeID       *string          `json:"knownMoveTypeId"`
	MinLevel              *int             `json:"minLevel"`
	MinHappiness          *int             `json:"minHappiness"`
	MinBeauty             *int             `json:"minBeauty"`
	MinAffection          *int             `json:"minAffection"`
	NeedsOverworldRain    bool             `json:"needsOverworldRain"`
	PartyPokemonID        *string          `json:"partyPokemonId"`
	PartyTypeID           *string          `json:"partyTypeId"`
	RelativePhysicalStats *int             `json:"relativePhysicalStats"`
	TimeOfDay             TimeOfDay        `json:"timeOfDay"`
	TradeWithPokemonID    *string          `json:"tradeWithPokemonId"`
	TurnUpsideDown        bool             `json:"turnUpsideDown"`
	Spin                  bool             `json:"spin"`
	TakeDamage            *int             `json:"takeDamage"`
	CriticalHits          *int             `json:"criticalHits"`
}

func (PokemonEvolution) IsNode() {}

func NewPokemonEvolutionFromDb(dbPokemonEvolution db.PokemonEvolutionModel) PokemonEvolution {
	pe := PokemonEvolution{
		FromPokemonID:      dbPokemonEvolution.FromPokemonID,
		ToPokemonID:        dbPokemonEvolution.ToPokemonID,
		Trigger:            EvolutionTrigger(dbPokemonEvolution.Trigger),
		Gender:             Gender(dbPokemonEvolution.Gender),
		NeedsOverworldRain: dbPokemonEvolution.NeedsOverworldRain,
		TimeOfDay:          TimeOfDay(dbPokemonEvolution.TimeOfDay),
		TurnUpsideDown:     dbPokemonEvolution.TurnUpsideDown,
		Spin:               dbPokemonEvolution.Spin,
	}

	if itemId, ok := dbPokemonEvolution.ItemID(); ok {
		pe.ItemID = &itemId
	} else {
		pe.ItemID = nil
	}

	if heldItemId, ok := dbPokemonEvolution.HeldItemID(); ok {
		pe.HeldItemID = &heldItemId
	} else {
		pe.HeldItemID = nil
	}

	if knownMoveId, ok := dbPokemonEvolution.KnownMoveID(); ok {
		pe.KnownMoveID = &knownMoveId
	} else {
		pe.KnownMoveID = nil
	}

	if knownMoveTypeId, ok := dbPokemonEvolution.KnownMoveTypeID(); ok {
		pe.KnownMoveTypeID = &knownMoveTypeId
	} else {
		pe.KnownMoveTypeID = nil
	}

	if minLevel, ok := dbPokemonEvolution.MinLevel(); ok {
		pe.MinLevel = &minLevel
	} else {
		pe.MinLevel = nil
	}

	if minHappiness, ok := dbPokemonEvolution.MinHappiness(); ok {
		pe.MinHappiness = &minHappiness
	} else {
		pe.MinHappiness = nil
	}

	if minBeauty, ok := dbPokemonEvolution.MinBeauty(); ok {
		pe.MinBeauty = &minBeauty
	} else {
		pe.MinBeauty = nil
	}

	if minAffection, ok := dbPokemonEvolution.MinAffection(); ok {
		pe.MinAffection = &minAffection
	} else {
		pe.MinAffection = nil
	}

	if partyPokemonId, ok := dbPokemonEvolution.PartyPokemonID(); ok {
		pe.PartyPokemonID = &partyPokemonId
	} else {
		pe.PartyPokemonID = nil
	}

	if partyTypeId, ok := dbPokemonEvolution.PartyTypeID(); ok {
		pe.PartyTypeID = &partyTypeId
	} else {
		pe.PartyTypeID = nil
	}

	if relativePhysicalStats, ok := dbPokemonEvolution.RelativePhysicalStats(); ok {
		pe.RelativePhysicalStats = &relativePhysicalStats
	} else {
		pe.RelativePhysicalStats = nil
	}

	if tradeWithPokemonId, ok := dbPokemonEvolution.TradeWithPokemonID(); ok {
		pe.TradeWithPokemonID = &tradeWithPokemonId
	} else {
		pe.TradeWithPokemonID = nil
	}

	if takeDamage, ok := dbPokemonEvolution.TakeDamage(); ok {
		pe.TakeDamage = &takeDamage
	} else {
		pe.TakeDamage = nil
	}

	if criticalHits, ok := dbPokemonEvolution.CriticalHits(); ok {
		pe.CriticalHits = &criticalHits
	} else {
		pe.CriticalHits = nil
	}

	return pe
}

func NewPokemonEvolutionEdgeFromDb(dbPokemonEvolution db.PokemonEvolutionModel) PokemonEvolutionEdge {
	node := NewPokemonEvolutionFromDb(dbPokemonEvolution)
	return PokemonEvolutionEdge{
		Cursor: dbPokemonEvolution.ID,
		Node:   &node,
	}
}

func NewEmptyPokemonEvolutionConnection() PokemonEvolutionConnection {
	return PokemonEvolutionConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonEvolutionEdge{},
	}
}

func (c *PokemonEvolutionConnection) AddEdge(e *PokemonEvolutionEdge) {
	if c.PageInfo.StartCursor == nil {
		c.PageInfo.StartCursor = &e.Cursor
	}
	c.PageInfo.EndCursor = &e.Cursor
	c.Edges = append(c.Edges, e)
}
