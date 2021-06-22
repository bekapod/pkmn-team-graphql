package model

import (
	"bekapod/pkmn-team-graphql/data/db"
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestNewPokemonEvolutionFromDb_WithNulls(t *testing.T) {
	pokemonEvolution := db.PokemonEvolutionModel{
		InnerPokemonEvolution: db.InnerPokemonEvolution{
			FromPokemonID:      "pokemon-1",
			ToPokemonID:        "pokeemon-2",
			Trigger:            db.EvolutionTriggerUSEITEM,
			Gender:             db.GenderFEMALE,
			NeedsOverworldRain: false,
			TimeOfDay:          db.TimeOfDayANY,
			TurnUpsideDown:     true,
			Spin:               false,
		},
	}
	exp := PokemonEvolution{
		FromPokemonID:      pokemonEvolution.FromPokemonID,
		ToPokemonID:        pokemonEvolution.ToPokemonID,
		Trigger:            EvolutionTriggerUseItem,
		Gender:             GenderFemale,
		NeedsOverworldRain: pokemonEvolution.NeedsOverworldRain,
		TimeOfDay:          TimeOfDayAny,
		TurnUpsideDown:     pokemonEvolution.TurnUpsideDown,
		Spin:               pokemonEvolution.Spin,
	}

	got := NewPokemonEvolutionFromDb(pokemonEvolution)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonEvolutionFromDb_WithFullData(t *testing.T) {
	itemId := "item-id"
	heldItemId := "held-item-id"
	knownMoveId := "known-move-id"
	knownMoveTypeId := "known-move-type-id"
	minLevel := 32
	minHappiness := 100
	minBeauty := 20
	minAffection := 50
	partyPokemonId := "party-pokemon-id"
	partyTypeId := "party-type-id"
	relativePhysicalStats := -1
	tradeWithPokemonId := "trade-with-pokemon-id"
	takeDamage := 100
	criticalHits := 5

	pokemonEvolution := db.PokemonEvolutionModel{
		InnerPokemonEvolution: db.InnerPokemonEvolution{
			FromPokemonID:         "pokemon-1",
			ToPokemonID:           "pokemon-2",
			Trigger:               db.EvolutionTriggerUSEITEM,
			ItemID:                &itemId,
			Gender:                db.GenderFEMALE,
			HeldItemID:            &heldItemId,
			KnownMoveID:           &knownMoveId,
			KnownMoveTypeID:       &knownMoveTypeId,
			MinLevel:              &minLevel,
			MinHappiness:          &minHappiness,
			MinBeauty:             &minBeauty,
			MinAffection:          &minAffection,
			NeedsOverworldRain:    false,
			PartyPokemonID:        &partyPokemonId,
			PartyTypeID:           &partyTypeId,
			RelativePhysicalStats: &relativePhysicalStats,
			TimeOfDay:             db.TimeOfDayANY,
			TradeWithPokemonID:    &tradeWithPokemonId,
			TurnUpsideDown:        true,
			Spin:                  false,
			TakeDamage:            &takeDamage,
			CriticalHits:          &criticalHits,
		},
	}
	exp := PokemonEvolution{
		FromPokemonID:         pokemonEvolution.FromPokemonID,
		ToPokemonID:           pokemonEvolution.ToPokemonID,
		Trigger:               EvolutionTriggerUseItem,
		ItemID:                &itemId,
		Gender:                GenderFemale,
		HeldItemID:            &heldItemId,
		KnownMoveID:           &knownMoveId,
		KnownMoveTypeID:       &knownMoveTypeId,
		MinLevel:              &minLevel,
		MinHappiness:          &minHappiness,
		MinBeauty:             &minBeauty,
		MinAffection:          &minAffection,
		NeedsOverworldRain:    pokemonEvolution.NeedsOverworldRain,
		PartyPokemonID:        &partyPokemonId,
		PartyTypeID:           &partyTypeId,
		RelativePhysicalStats: &relativePhysicalStats,
		TimeOfDay:             TimeOfDayAny,
		TradeWithPokemonID:    &tradeWithPokemonId,
		TurnUpsideDown:        pokemonEvolution.TurnUpsideDown,
		Spin:                  pokemonEvolution.Spin,
		TakeDamage:            &takeDamage,
		CriticalHits:          &criticalHits,
	}

	got := NewPokemonEvolutionFromDb(pokemonEvolution)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewPokemonEvolutionEdgeFromDb(t *testing.T) {
	pokemonEvolution := db.PokemonEvolutionModel{
		InnerPokemonEvolution: db.InnerPokemonEvolution{
			ID:                 "123",
			FromPokemonID:      "pokemon-1",
			ToPokemonID:        "pokeemon-2",
			Trigger:            db.EvolutionTriggerUSEITEM,
			Gender:             db.GenderFEMALE,
			NeedsOverworldRain: false,
			TimeOfDay:          db.TimeOfDayANY,
			TurnUpsideDown:     true,
			Spin:               false,
		},
	}
	exp := PokemonEvolutionEdge{
		Cursor: pokemonEvolution.ID,
		Node: &PokemonEvolution{
			FromPokemonID:      pokemonEvolution.FromPokemonID,
			ToPokemonID:        pokemonEvolution.ToPokemonID,
			Trigger:            EvolutionTriggerUseItem,
			Gender:             GenderFemale,
			NeedsOverworldRain: pokemonEvolution.NeedsOverworldRain,
			TimeOfDay:          TimeOfDayAny,
			TurnUpsideDown:     pokemonEvolution.TurnUpsideDown,
			Spin:               pokemonEvolution.Spin,
		},
	}

	got := NewPokemonEvolutionEdgeFromDb(pokemonEvolution)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonEvolutionConnection(t *testing.T) {
	exp := PokemonEvolutionConnection{
		PageInfo: &PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		},
		Edges: []*PokemonEvolutionEdge{},
	}

	got := NewEmptyPokemonEvolutionConnection()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonEvolutionConnection_AddEdge(t *testing.T) {
	pokemonEvolutions := NewEmptyPokemonEvolutionConnection()
	pokemonEvolution1 := &PokemonEvolutionEdge{}
	pokemonEvolution2 := &PokemonEvolutionEdge{}
	pokemonEvolutions.AddEdge(pokemonEvolution1)
	pokemonEvolutions.AddEdge(pokemonEvolution2)

	if !reflect.DeepEqual([]*PokemonEvolutionEdge{pokemonEvolution1, pokemonEvolution2}, pokemonEvolutions.Edges) {
		t.Errorf("the pokemon evolutions added do not match")
	}
}
