package model

import (
	"encoding/json"
	"fmt"
)

type Evolution struct {
	Trigger               EvolutionTrigger `json:"evolutionTrigger"`
	Item                  *Item            `json:"item"`
	Gender                Gender           `json:"gender"`
	HeldItem              *Item            `json:"heldItem"`
	Location              *Location        `json:"location"`
	MinLevel              int              `json:"minLevel"`
	MinHappiness          int              `json:"minHappiness"`
	MinBeauty             int              `json:"minBeauty"`
	MinAffection          int              `json:"minAffection"`
	NeedsOverworldRain    bool             `json:"needsOverworldRain"`
	RelativePhysicalStats int              `json:"relativePhysicalStats"`
	TimeOfDay             TimeOfDay        `json:"timeOfDay"`
	TurnUpsideDown        bool             `json:"turnUpsideDown"`
	Spin                  bool             `json:"spin"`
	TakeDamage            int              `json:"takeDamage"`
	CriticalHits          int              `json:"criticalHits"`
	PokemonID             *string          `json:"pokemon_id"`
	PartySpeciesPokemonID *string          `json:"party_species_pokemon_id"`
	TradeSpeciesPokemonID *string          `json:"trade_species_pokemon_id"`
	KnownMoveID           *string          `json:"known_move_id"`
	KnownMoveTypeID       *string          `json:"known_move_type_id"`
	PartyTypeID           *string          `json:"party_type_id"`
}

type EvolutionList struct {
	Total      int         `json:"total"`
	Evolutions []Evolution `json:"evolutions"`
}

func NewEvolutionList(e []Evolution) EvolutionList {
	return EvolutionList{
		Total:      len(e),
		Evolutions: e,
	}
}

func NewEmptyEvolutionList() EvolutionList {
	return EvolutionList{
		Total:      0,
		Evolutions: []Evolution{},
	}
}

func (l *EvolutionList) AddEvolution(e Evolution) {
	l.Total++
	l.Evolutions = append(l.Evolutions, e)
}

func (r *Evolution) Scan(src interface{}) error {
	switch v := src.(type) {
	case []uint8:
		err := json.Unmarshal([]byte(v), &r)
		return err
	}

	return fmt.Errorf("failed to scan evolution")
}

func (Evolution) IsEntity() {}
