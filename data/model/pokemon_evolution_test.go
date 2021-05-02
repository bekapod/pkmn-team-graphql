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

func TestNewPokemonEvolutionFromDb_WithFullDate(t *testing.T) {
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

func TestNewPokemonEvolutionList(t *testing.T) {
	pokemonEvolutions := []*PokemonEvolution{
		{
			FromPokemonID: "pokemon-1",
		},
		{
			FromPokemonID: "pokemon-1",
		},
	}

	exp := PokemonEvolutionList{
		Total:             2,
		PokemonEvolutions: pokemonEvolutions,
	}

	got := NewPokemonEvolutionList(pokemonEvolutions)
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestNewEmptyPokemonEvolutionList(t *testing.T) {
	exp := PokemonEvolutionList{
		Total:             0,
		PokemonEvolutions: []*PokemonEvolution{},
	}

	got := NewEmptyPokemonEvolutionList()
	if diff := deep.Equal(exp, got); diff != nil {
		t.Error(diff)
	}
}

func TestPokemonEvolutionList_AddPokemonEvolution(t *testing.T) {
	pokemonEvolutions := PokemonEvolutionList{}
	pokemonEvolution1 := &PokemonEvolution{}
	pokemonEvolution2 := &PokemonEvolution{}
	pokemonEvolutions.AddPokemonEvolution(pokemonEvolution1)
	pokemonEvolutions.AddPokemonEvolution(pokemonEvolution2)

	if pokemonEvolutions.Total != 2 {
		t.Errorf("expected Total of 2, but got %d instead", pokemonEvolutions.Total)
	}

	if !reflect.DeepEqual([]*PokemonEvolution{pokemonEvolution1, pokemonEvolution2}, pokemonEvolutions.PokemonEvolutions) {
		t.Errorf("the pokemon evolutions added do not match")
	}
}
