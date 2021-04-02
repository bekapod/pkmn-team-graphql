package pokeapi

import (
	"bekapod/pkmn-team-graphql/log"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PokeApiConfig struct {
	Host   string `arg:"--pokeapi-host,env:POKEAPI_HOST" default:"https://pokeapi.co"`
	Prefix string `default:"api/v2"`
}

type PokeApiClient struct {
	httpClient *http.Client
	config     *PokeApiConfig
}

func New(config PokeApiConfig) *PokeApiClient {
	return &PokeApiClient{
		httpClient: &http.Client{Timeout: 60 * time.Second},
		config:     &config,
	}
}

func (client PokeApiClient) GetResourceList(resource string) resourcePointerList {
	result := resourcePointerList{}
	client.getResource(resource+"?limit=10000", &result)
	return result
}

func (client PokeApiClient) GetAbility(idOrName string) ability {
	result := ability{}
	client.getResource(fmt.Sprintf("ability/%s", idOrName), &result)
	return result
}

func (client PokeApiClient) GetEggGroup(idOrName string) eggGroup {
	result := eggGroup{}
	client.getResource(fmt.Sprintf("egg-group/%s", idOrName), &result)
	return result
}

func (client PokeApiClient) GetMove(idOrName string) move {
	result := move{}
	client.getResource(fmt.Sprintf("move/%s", idOrName), &result)
	return result
}

func (client PokeApiClient) GetMoveTarget(nameOrId string) target {
	result := target{}
	client.getResource(fmt.Sprintf("move-target/%s", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetPokemon(nameOrId string) pokemon {
	result := pokemon{}
	client.getResource(fmt.Sprintf("pokemon/%s", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetPokemonSpecies(nameOrId string) pokemonSpecies {
	result := pokemonSpecies{}
	client.getResource(fmt.Sprintf("pokemon-species/%s", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetPokemonVarietiesForSpecies(species pokemonSpecies) []pokemon {
	pokemon := make([]pokemon, 0)

	for i := range species.Varieties {
		urlParts := strings.Split(species.Varieties[i].Pokemon.Url, "/")
		id := urlParts[len(urlParts)-2]
		variety := client.GetPokemon(id)
		if !strings.Contains(variety.Name, "gmax") && !strings.Contains(variety.Name, "totem") && !strings.Contains(variety.Name, "-mega") {
			pokemon = append(pokemon, variety)
		}
	}

	return pokemon
}

func (client PokeApiClient) GetStat(nameOrId string) stat {
	result := stat{}
	client.getResource(fmt.Sprintf("stat/%s", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetType(nameOrId string) pokeType {
	result := pokeType{}
	client.getResource(fmt.Sprintf("type/%s", nameOrId), &result)
	return result
}

func (client PokeApiClient) getResource(path string, bucket interface{}) {
	resourceUrl := fmt.Sprintf("%s/%s/%s", client.config.Host, client.config.Prefix, path)
	log.Logger.WithField("resource", resourceUrl).Info("Getting resource")
	response, err := client.httpClient.Get(resourceUrl)
	check(err)
	log.Logger.WithField("resource", resourceUrl).Info("Got resource")
	defer response.Body.Close()
	decodeErr := json.NewDecoder(response.Body).Decode(bucket)
	check(decodeErr)
}

func GetEnglishName(names []*translatedName, resourceName string) (*translatedName, error) {
	for i := range names {
		if names[i].Language.Name == "en" {
			return names[i], nil
		}
	}

	return &translatedName{}, fmt.Errorf("no english name found for %s", resourceName)
}

func GetEnglishGenus(genera []*genus, resourceName string) (*genus, error) {
	for i := range genera {
		if genera[i].Language.Name == "en" {
			return genera[i], nil
		}
	}

	return &genus{
		Genus: "",
	}, fmt.Errorf("no english genus found for %s", resourceName)
}

func GetEnglishEffectEntry(effectEntries []*effectEntry, resourceName string) (*effectEntry, error) {
	for i := range effectEntries {
		if effectEntries[i].Language.Name == "en" {
			return effectEntries[i], nil
		}
	}

	return &effectEntry{
		Effect:      "",
		ShortEffect: "",
	}, fmt.Errorf("no english effect entry found for %s", resourceName)
}

func GetEnglishFlavourTextEntry(flavourTextEntries []*flavourTextEntry, resourceName string) (*flavourTextEntry, error) {
	for i := range flavourTextEntries {
		if flavourTextEntries[i].Language.Name == "en" && (flavourTextEntries[i].VersionGroup.Name == "ultra-sun-ultra-moon" || flavourTextEntries[i].Version.Name == "ultra-moon" || flavourTextEntries[i].Version.Name == "lets-go-pikachu") {
			return flavourTextEntries[i], nil
		}
	}

	return &flavourTextEntry{
		FlavourText: "",
	}, fmt.Errorf("no english flavour text entry found for %s", resourceName)
}

func GetPokemonStat(pkmn pokemon, stat string) (int, error) {
	for i := range pkmn.Stats {
		if pkmn.Stats[i].Stat.Name == stat {
			return pkmn.Stats[i].BaseStat, nil
		}
	}

	return 0, fmt.Errorf("couldn't find %s stat for %s", stat, pkmn.Name)
}

func GetPokedexId(pkmn pokemonSpecies, pokedex string) (int, error) {
	for i := range pkmn.PokedexNumbers {
		if pkmn.PokedexNumbers[i].Pokedex.Name == pokedex {
			return pkmn.PokedexNumbers[i].EntryNumber, nil
		}
	}

	return 0, fmt.Errorf("couldn't find %s pokedex id for %s", pokedex, pkmn.Name)
}

func check(e error) {
	if e != nil {
		log.Logger.Fatal(e)
	}
}

type resourcePointerList struct {
	Count   int                `json:"count"`
	Results []*resourcePointer `json:"results"`
}

type resourcePointer struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type translatedName struct {
	Name     string          `json:"name"`
	Language resourcePointer `json:"language"`
}

type effectEntry struct {
	Effect      string          `json:"effect"`
	ShortEffect string          `json:"short_effect"`
	Language    resourcePointer `json:"language"`
}

type flavourTextEntry struct {
	FlavourText  string          `json:"flavor_text"`
	Language     resourcePointer `json:"language"`
	VersionGroup resourcePointer `json:"version_group"`
	Version      resourcePointer `json:"version"`
}

type pokeType struct {
	Id              int                `json:"id"`
	Name            string             `json:"name"`
	Names           []*translatedName  `json:"names"`
	Moves           []*resourcePointer `json:"moves"`
	DamageRelations struct {
		NoDamageTo       []*resourcePointer `json:"no_damage_to"`
		HalfDamageTo     []*resourcePointer `json:"half_damage_to"`
		DoubleDamageTo   []*resourcePointer `json:"double_damage_to"`
		NoDamageFrom     []*resourcePointer `json:"no_damage_from"`
		HalfDamageFrom   []*resourcePointer `json:"half_damage_from"`
		DoubleDamageFrom []*resourcePointer `json:"double_damage_from"`
	} `json:"damage_relations"`
	Pokemon []*struct {
		Slot    int             `json:"slot"`
		Pokemon resourcePointer `json:"pokemon"`
	} `json:"pokemon"`
}

type move struct {
	Id            int                 `json:"id"`
	Name          string              `json:"name"`
	Accuracy      int                 `json:"accuracy"`
	EffectChance  int                 `json:"effect_chance"`
	PP            int                 `json:"pp"`
	Priority      int                 `json:"priority"`
	Power         int                 `json:"power"`
	DamageClass   resourcePointer     `json:"damage_class"`
	EffectEntries []*effectEntry      `json:"effect_entries"`
	Names         []*translatedName   `json:"names"`
	Target        resourcePointer     `json:"target"`
	Type          resourcePointer     `json:"type"`
	Description   []*flavourTextEntry `json:"flavor_text_entries"`
}

type target struct {
	Id    int               `json:"id"`
	Name  string            `json:"name"`
	Names []*translatedName `json:"names"`
}

type genus struct {
	Genus    string          `json:"genus"`
	Language resourcePointer `json:"language"`
}

type ability struct {
	Id                 int                 `json:"id"`
	Name               string              `json:"name"`
	IsMainSeries       bool                `json:"is_main_series"`
	Names              []*translatedName   `json:"names"`
	EffectEntries      []*effectEntry      `json:"effect_entries"`
	FlavourTextEntries []*flavourTextEntry `json:"flavor_text_entries"`
	Pokemon            []*struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Pokemon  resourcePointer
	} `json:"pokemon"`
}

type stat struct {
	ID           int               `json:"id"`
	Name         string            `json:"name"`
	GameIndex    int               `json:"game_index"`
	IsBattleOnly bool              `json:"is_battle_only"`
	Names        []*translatedName `json:"names"`
}

type eggGroup struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Names          []*translatedName  `json:"names"`
	PokemonSpecies []*resourcePointer `json:"pokemon_species"`
}

type pokemon struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	BaseExperience int                `json:"base_experience"`
	Height         int                `json:"height"`
	IsDefault      bool               `json:"is_default"`
	Order          int                `json:"order"`
	Weight         int                `json:"weight"`
	Forms          []*resourcePointer `json:"forms"`
	Species        resourcePointer    `json:"species"`
	Abilities      []*struct {
		IsHidden bool            `json:"is_hidden"`
		Slot     int             `json:"slot"`
		Ability  resourcePointer `json:"ability"`
	} `json:"abilities"`
	Moves []*struct {
		Move                resourcePointer `json:"move"`
		VersionGroupDetails []*struct {
			LevelLearnedAt  int             `json:"level_learned_at"`
			VersionGroup    resourcePointer `json:"version_group"`
			MoveLearnMethod resourcePointer `json:"move_learn_method"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Sprites struct {
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  string `json:"front_female"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
			} `json:"official_artwork"`
		} `json:"other"`
	} `json:"sprites"`
	Stats []*struct {
		BaseStat int             `json:"base_stat"`
		Effort   int             `json:"effort"`
		Stat     resourcePointer `json:"stat"`
	} `json:"stats"`
	Types []*struct {
		Slot int             `json:"slot"`
		Type resourcePointer `json:"type"`
	} `json:"types"`
}

type pokemonSpecies struct {
	ID                   int                 `json:"id"`
	Name                 string              `json:"name"`
	Order                int                 `json:"order"`
	GenderRate           int                 `json:"gender_rate"`
	CaptureRate          int                 `json:"capture_rate"`
	BaseHappiness        int                 `json:"base_happiness"`
	IsBaby               bool                `json:"is_baby"`
	IsLegendary          bool                `json:"is_legendary"`
	IsMythical           bool                `json:"is_mythical"`
	HatchCounter         int                 `json:"hatch_counter"`
	HasGenderDifferences bool                `json:"has_gender_differences"`
	FormsSwitchable      bool                `json:"forms_switchable"`
	GrowthRate           resourcePointer     `json:"growth_rate"`
	EggGroups            []*resourcePointer  `json:"egg_groups"`
	Color                resourcePointer     `json:"color"`
	Shape                resourcePointer     `json:"shape"`
	EvolvesFromSpecies   resourcePointer     `json:"evolves_from_species"`
	EvolutionChain       resourcePointer     `json:"evolution_chain"`
	Habitat              resourcePointer     `json:"habitat"`
	Names                []*translatedName   `json:"names"`
	FlavourTextEntries   []*flavourTextEntry `json:"flavor_text_entries"`
	Genera               []*genus            `json:"genera"`
	PokedexNumbers       []*struct {
		EntryNumber int             `json:"entry_number"`
		Pokedex     resourcePointer `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Varieties []*struct {
		IsDefault bool            `json:"is_default"`
		Pokemon   resourcePointer `json:"pokemon"`
	} `json:"varieties"`
}
