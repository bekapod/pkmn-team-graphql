package pokeapi

import (
	"fmt"
)

type ResourcePointerList struct {
	Count   int                `json:"count"`
	Results []*ResourcePointer `json:"results"`
}

type ResourcePointer struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type RawTranslatedName struct {
	Name     string          `json:"name"`
	Language ResourcePointer `json:"language"`
}

type RawDamageRelations struct {
	NoDamageTo       []*ResourcePointer `json:"no_damage_to"`
	HalfDamageTo     []*ResourcePointer `json:"half_damage_to"`
	DoubleDamageTo   []*ResourcePointer `json:"double_damage_to"`
	NoDamageFrom     []*ResourcePointer `json:"no_damage_from"`
	HalfDamageFrom   []*ResourcePointer `json:"half_damage_from"`
	DoubleDamageFrom []*ResourcePointer `json:"double_damage_from"`
}

type RawPokemonWithType struct {
	Slot    int             `json:"slot"`
	Pokemon ResourcePointer `json:"pokemon"`
}

type RawEffectEntry struct {
	Effect      string          `json:"effect"`
	ShortEffect string          `json:"short_effect"`
	Language    ResourcePointer `json:"language"`
}

type RawFlavour struct {
	FlavourText  string          `json:"flavor_text"`
	Language     ResourcePointer `json:"language"`
	VersionGroup ResourcePointer `json:"version_group"`
}

type RawType struct {
	Id              int                   `json:"id"`
	Name            string                `json:"name"`
	DamageRelations RawDamageRelations    `json:"damage_relations"`
	Names           []*RawTranslatedName  `json:"names"`
	Pokemon         []*RawPokemonWithType `json:"pokemon"`
	Moves           []*ResourcePointer    `json:"moves"`
}

type RawMove struct {
	Id            int                  `json:"id"`
	Name          string               `json:"name"`
	Accuracy      int                  `json:"accuracy"`
	EffectChance  int                  `json:"effect_chance"`
	PP            int                  `json:"pp"`
	Priority      int                  `json:"priority"`
	Power         int                  `json:"power"`
	DamageClass   ResourcePointer      `json:"damage_class"`
	EffectEntries []*RawEffectEntry    `json:"effect_entries"`
	Names         []*RawTranslatedName `json:"names"`
	Target        ResourcePointer      `json:"target"`
	Type          ResourcePointer      `json:"type"`
	Description   []*RawFlavour        `json:"flavor_text_entries"`
}

type RawTarget struct {
	Id    int                  `json:"id"`
	Name  string               `json:"name"`
	Names []*RawTranslatedName `json:"names"`
}

type RawAbility struct {
	Id                 int                  `json:"id"`
	Name               string               `json:"name"`
	IsMainSeries       bool                 `json:"is_main_series"`
	Names              []*RawTranslatedName `json:"names"`
	EffectEntries      []*RawEffectEntry    `json:"effect_entries"`
	FlavourTextEntries []*RawFlavour        `json:"flavor_text_entries"`
}

func GetEnglishName(names []*RawTranslatedName, resourceName string) (*RawTranslatedName, error) {
	for i := range names {
		if names[i].Language.Name == "en" {
			return names[i], nil
		}
	}

	return &RawTranslatedName{}, fmt.Errorf("no english name found for %s", resourceName)
}

func GetEnglishEffectEntry(effectEntries []*RawEffectEntry, resourceName string) (*RawEffectEntry, error) {
	for i := range effectEntries {
		if effectEntries[i].Language.Name == "en" {
			return effectEntries[i], nil
		}
	}

	return &RawEffectEntry{
		Effect:      "",
		ShortEffect: "",
	}, fmt.Errorf("no english effect entry found for %s", resourceName)
}
