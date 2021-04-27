package model

import (
	"encoding/json"
	"fmt"
)

type Evolution struct {
	Pokemon               Pokemon          `json:"pokemon"`
	Trigger               EvolutionTrigger `json:"evolutionTrigger"`
	Item                  Item             `json:"item"`
	Gender                Gender           `json:"gender"`
	HeldItem              Item             `json:"heldItem"`
	KnownMove             Move             `json:"knownMove"`
	Location              Location         `json:"location"`
	MinLevel              int              `json:"minLevel"`
	MinHappiness          int              `json:"minHappiness"`
	MinBeauty             int              `json:"minBeauty"`
	MinAffection          int              `json:"minAffection"`
	NeedsOverworldRain    bool             `json:"needsOverworldRain"`
	PartyPokemon          Pokemon          `json:"partyPokemon"`
	RelativePhysicalStats int              `json:"relativePhysicalStats"`
	TimeOfDay             TimeOfDay        `json:"timeOfDay"`
	TradeWithPokemon      Pokemon          `json:"tradeWithPokemon"`
	TurnUpsideDown        bool             `json:"turnUpsideDown"`
	Spin                  bool             `json:"spin"`
	TakeDamage            int              `json:"takeDamage"`
	CriticalHits          int              `json:"criticalHits"`
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
