package main

import (
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/pokeapi"
	"bekapod/pkmn-team-graphql/scripts/seed/helpers"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
)

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	start := time.Now()
	config := &helpers.Config{
		OutputFile: "seeds/pokemon.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)

	generation := client.GetGeneration("generation-viii")
	pokemonSpeciesList := make([]string, 0)
	for _, versionGroupResource := range generation.VersionGroups {
		versionGroup := client.GetVersionGroup(versionGroupResource.Name)

		for _, pokedexResource := range versionGroup.Pokedexes {
			pokedex := client.GetPokedex(pokedexResource.Name)

			for _, pokemonEntry := range pokedex.PokemonEntries {
				pokemonSpeciesList = append(pokemonSpeciesList, pokemonEntry.PokemonSpecies.Url)
			}
		}
	}

	otherAvailablePokemon := []string{
		"http://localhost/api/v2/pokemon-species/1/",
		"http://localhost/api/v2/pokemon-species/2/",
		"http://localhost/api/v2/pokemon-species/3/",
		"http://localhost/api/v2/pokemon-species/7/",
		"http://localhost/api/v2/pokemon-species/8/",
		"http://localhost/api/v2/pokemon-species/9/",
		"http://localhost/api/v2/pokemon-species/150/",
		"http://localhost/api/v2/pokemon-species/151/",
		"http://localhost/api/v2/pokemon-species/240/",
		"http://localhost/api/v2/pokemon-species/243/",
		"http://localhost/api/v2/pokemon-species/244/",
		"http://localhost/api/v2/pokemon-species/245/",
		"http://localhost/api/v2/pokemon-species/249/",
		"http://localhost/api/v2/pokemon-species/250/",
		"http://localhost/api/v2/pokemon-species/252/",
		"http://localhost/api/v2/pokemon-species/253/",
		"http://localhost/api/v2/pokemon-species/254/",
		"http://localhost/api/v2/pokemon-species/255/",
		"http://localhost/api/v2/pokemon-species/256/",
		"http://localhost/api/v2/pokemon-species/257/",
		"http://localhost/api/v2/pokemon-species/258/",
		"http://localhost/api/v2/pokemon-species/259/",
		"http://localhost/api/v2/pokemon-species/260/",
		"http://localhost/api/v2/pokemon-species/276/",
		"http://localhost/api/v2/pokemon-species/380/",
		"http://localhost/api/v2/pokemon-species/381/",
		"http://localhost/api/v2/pokemon-species/382/",
		"http://localhost/api/v2/pokemon-species/383/",
		"http://localhost/api/v2/pokemon-species/384/",
		"http://localhost/api/v2/pokemon-species/385/",
		"http://localhost/api/v2/pokemon-species/480/",
		"http://localhost/api/v2/pokemon-species/481/",
		"http://localhost/api/v2/pokemon-species/482/",
		"http://localhost/api/v2/pokemon-species/483/",
		"http://localhost/api/v2/pokemon-species/484/",
		"http://localhost/api/v2/pokemon-species/485/",
		"http://localhost/api/v2/pokemon-species/486/",
		"http://localhost/api/v2/pokemon-species/487/",
		"http://localhost/api/v2/pokemon-species/488/",
		"http://localhost/api/v2/pokemon-species/641/",
		"http://localhost/api/v2/pokemon-species/642/",
		"http://localhost/api/v2/pokemon-species/643/",
		"http://localhost/api/v2/pokemon-species/644/",
		"http://localhost/api/v2/pokemon-species/645/",
		"http://localhost/api/v2/pokemon-species/646/",
		"http://localhost/api/v2/pokemon-species/647/",
		"http://localhost/api/v2/pokemon-species/716/",
		"http://localhost/api/v2/pokemon-species/717/",
		"http://localhost/api/v2/pokemon-species/718/",
		"http://localhost/api/v2/pokemon-species/722/",
		"http://localhost/api/v2/pokemon-species/723/",
		"http://localhost/api/v2/pokemon-species/724/",
		"http://localhost/api/v2/pokemon-species/725/",
		"http://localhost/api/v2/pokemon-species/726/",
		"http://localhost/api/v2/pokemon-species/727/",
		"http://localhost/api/v2/pokemon-species/728/",
		"http://localhost/api/v2/pokemon-species/729/",
		"http://localhost/api/v2/pokemon-species/730/",
		"http://localhost/api/v2/pokemon-species/785/",
		"http://localhost/api/v2/pokemon-species/786/",
		"http://localhost/api/v2/pokemon-species/787/",
		"http://localhost/api/v2/pokemon-species/788/",
		"http://localhost/api/v2/pokemon-species/789/",
		"http://localhost/api/v2/pokemon-species/790/",
		"http://localhost/api/v2/pokemon-species/791/",
		"http://localhost/api/v2/pokemon-species/792/",
		"http://localhost/api/v2/pokemon-species/793/",
		"http://localhost/api/v2/pokemon-species/794/",
		"http://localhost/api/v2/pokemon-species/795/",
		"http://localhost/api/v2/pokemon-species/796/",
		"http://localhost/api/v2/pokemon-species/797/",
		"http://localhost/api/v2/pokemon-species/798/",
		"http://localhost/api/v2/pokemon-species/799/",
		"http://localhost/api/v2/pokemon-species/800/",
		"http://localhost/api/v2/pokemon-species/803/",
		"http://localhost/api/v2/pokemon-species/804/",
		"http://localhost/api/v2/pokemon-species/805/",
		"http://localhost/api/v2/pokemon-species/806/",
		"http://localhost/api/v2/pokemon-species/807/",
		"http://localhost/api/v2/pokemon-species/808/",
		"http://localhost/api/v2/pokemon-species/809/",
	}
	pokemonSpeciesList = append(pokemonSpeciesList, otherAvailablePokemon...)

	pokemonSpeciesList = unique(pokemonSpeciesList)
	resultsLength := len(pokemonSpeciesList)
	pokemonValues := make([]string, 0)
	pokemonTypeValues := make([]string, 0)
	pokemonAbilityValues := make([]string, 0)
	pokemonMoveValues := make([]string, 0)
	pokemonEggGroupValues := make([]string, 0)
	pokemonEvolutionValues := make([]string, 0)
	pokemonEvolutionChainList := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)
	sem := make(chan bool, 20)

	for i := 0; i < resultsLength; i++ {
		sem <- true
		go func(i int) {
			defer func() { <-sem }()
			defer wg.Done()
			urlParts := strings.Split(pokemonSpeciesList[i], "/")
			id := urlParts[len(urlParts)-2]
			fullPokemonSpecies := client.GetPokemonSpecies(id)
			varieties := client.GetPokemonVarietiesForSpecies(fullPokemonSpecies)

			pokedexId, err := pokeapi.GetPokedexId(fullPokemonSpecies, "national")
			if err != nil {
				log.Logger.Fatal(err)
			}
			englishName, err := pokeapi.GetEnglishName(fullPokemonSpecies.Names, fullPokemonSpecies.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			englishFlavourText, _ := pokeapi.GetEnglishFlavourTextEntry(fullPokemonSpecies.FlavourTextEntries, fullPokemonSpecies.Name)
			englishGenus, err := pokeapi.GetEnglishGenus(fullPokemonSpecies.Genera, fullPokemonSpecies.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			habitat := "NULL"
			if fullPokemonSpecies.Habitat.Name != "" {
				habitat = fmt.Sprintf("'%s'", fullPokemonSpecies.Habitat.Name)
			}

			for i := range varieties {
				pokemon := varieties[i]
				hp, _ := pokeapi.GetPokemonStat(pokemon, "hp")
				attack, _ := pokeapi.GetPokemonStat(pokemon, "attack")
				defense, _ := pokeapi.GetPokemonStat(pokemon, "defense")
				specialAttack, _ := pokeapi.GetPokemonStat(pokemon, "special-attack")
				specialDefense, _ := pokeapi.GetPokemonStat(pokemon, "special-defense")
				speed, _ := pokeapi.GetPokemonStat(pokemon, "speed")

				pokemonValues = append(pokemonValues, fmt.Sprintf(
					"(%d, '%s', '%s', '%s', %d, %d, %d, %d, %d, %d, %t, %t, %t, '%s', '%s', '%s', %s, %t, '%s', %d, %d, now())",
					pokedexId,
					pokemon.Name,
					helpers.EscapeSingleQuote(englishName.Name),
					pokemon.Sprites.FrontDefault,
					hp,
					attack,
					defense,
					specialAttack,
					specialDefense,
					speed,
					fullPokemonSpecies.IsBaby,
					fullPokemonSpecies.IsLegendary,
					fullPokemonSpecies.IsMythical,
					helpers.EscapeSingleQuote(englishFlavourText.FlavourText),
					fullPokemonSpecies.Color.Name,
					fullPokemonSpecies.Shape.Name,
					habitat,
					pokemon.IsDefault,
					englishGenus.Genus,
					pokemon.Height,
					pokemon.Weight,
				))

				for i := range pokemon.Types {
					pokemonTypeValues = append(pokemonTypeValues, fmt.Sprintf(
						"(%s, %s, %d, now())",
						fmt.Sprintf("(SELECT id from pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id from types WHERE slug='%s')", pokemon.Types[i].Type.Name),
						pokemon.Types[i].Slot,
					))
				}

				for i := range pokemon.Abilities {
					abilityName := pokemon.Abilities[i].Ability.Name

					if abilityName == "as-one" && pokemon.Name == "calyrex-shadow-rider" {
						abilityName = "as-one-shadow-rider"
					}

					if abilityName == "as-one" && pokemon.Name == "calyrex-ice-rider" {
						abilityName = "as-one-ice-rider"
					}

					pokemonAbilityValues = append(pokemonAbilityValues, fmt.Sprintf(
						"(%s, %s, %d, %t, now())",
						fmt.Sprintf("(SELECT id FROM pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id FROM abilities WHERE slug='%s')", abilityName),
						pokemon.Abilities[i].Slot,
						pokemon.Abilities[i].IsHidden,
					))
				}

				for _, move := range pokemon.Moves {
					for _, versionGroup := range move.VersionGroupDetails {
						if versionGroup.VersionGroup.Name == "sword-shield" && !strings.Contains(move.Move.Name, "max-") {
							pokemonMoveValues = append(pokemonMoveValues, fmt.Sprintf(
								"(%s, %s)",
								fmt.Sprintf("(SELECT id FROM pokemon WHERE slug='%s')", pokemon.Name),
								fmt.Sprintf("(SELECT id FROM moves WHERE slug='%s')", move.Move.Name),
							))
						}
					}
				}

				for i := range fullPokemonSpecies.EggGroups {
					pokemonEggGroupValues = append(pokemonEggGroupValues, fmt.Sprintf(
						"(%s, %s)",
						fmt.Sprintf("(SELECT id FROM pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id FROM egg_groups WHERE slug='%s')", fullPokemonSpecies.EggGroups[i].Name),
					))
				}

				evolutionUrlParts := strings.Split(fullPokemonSpecies.EvolutionChain.Url, "/")
				evolutionChainId := evolutionUrlParts[len(evolutionUrlParts)-2]
				pokemonEvolutionChainList = append(pokemonEvolutionChainList, evolutionChainId)
			}
		}(i)
	}

	wg.Wait()

	pokemonEvolutionChainList = unique(pokemonEvolutionChainList)
	pokemonEvolutionChainListLength := len(pokemonEvolutionChainList)
	var wg2 sync.WaitGroup
	wg2.Add(pokemonEvolutionChainListLength)
	sem2 := make(chan bool, 20)

	for i := 0; i < pokemonEvolutionChainListLength; i++ {
		sem2 <- true
		go func(i int) {
			defer func() { <-sem2 }()
			defer wg2.Done()

			evolutionChain := client.GetEvolutionChain(pokemonEvolutionChainList[i])
			for _, chain := range evolutionChain.Chain.EvolvesTo {
				evolutionValues := parseEvolution(evolutionChain.Chain, chain)
				pokemonEvolutionValues = append(pokemonEvolutionValues, evolutionValues...)
			}
		}(i)
	}

	sql := fmt.Sprintf(
		"INSERT INTO pokemon (pokedex_id, slug, name, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, shape_enum, habitat_enum, is_default_variant, genus, height, weight, updated_at)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tpokedex_id = EXCLUDED.pokedex_id,\n\t\tname = EXCLUDED.name,\n\t\tsprite = EXCLUDED.sprite,\n\t\thp = EXCLUDED.hp,\n\t\tattack = EXCLUDED.attack,\n\t\tdefense = EXCLUDED.defense,\n\t\tspecial_attack = EXCLUDED.special_attack,\n\t\tspecial_defense = EXCLUDED.special_defense,\n\t\tspeed = EXCLUDED.speed,\n\t\tis_baby = EXCLUDED.is_baby,\n\t\tis_legendary = EXCLUDED.is_legendary,\n\t\tis_mythical = EXCLUDED.is_mythical,\n\t\tdescription = EXCLUDED.description,\n\t\tcolor_enum = EXCLUDED.color_enum,\n\t\tshape_enum = EXCLUDED.shape_enum,\n\t\thabitat_enum = EXCLUDED.habitat_enum,\n\t\tis_default_variant = EXCLUDED.is_default_variant,\n\t\tgenus = EXCLUDED.genus,\n\t\theight = EXCLUDED.height,\n\t\tweight = EXCLUDED.weight,\n\t\tupdated_at = EXCLUDED.updated_at;\n\n"+
			"INSERT INTO pokemon_type (pokemon_id, type_id, slot, updated_at)\n\tVALUES %s\nON CONFLICT (pokemon_id, type_id)\n\tDO UPDATE SET\n\t\tslot = EXCLUDED.slot,\n\t\tupdated_at = EXCLUDED.UPDATED_AT;\n\n"+
			"INSERT INTO pokemon_ability (pokemon_id, ability_id, slot, is_hidden, updated_at)\n\tVALUES %s\nON CONFLICT (pokemon_id, ability_id)\n\tDO UPDATE SET\n\t\tslot = EXCLUDED.slot, is_hidden = EXCLUDED.is_hidden, updated_at = EXCLUDED.updated_at;\n\n"+
			"INSERT INTO pokemon_egg_group (pokemon_id, egg_group_id)\n\tVALUES %s\nON CONFLICT (pokemon_id, egg_group_id)\n\tDO NOTHING;"+
			"INSERT INTO pokemon_evolutions (from_pokemon_id, to_pokemon_id, trigger_enum, item_id, gender_enum, held_item_id, known_move_id, known_move_type_id, location_id, min_level, min_happiness, min_beauty, min_affection, needs_overworld_rain, party_species_pokemon_id, party_type_id, relative_physical_stats, time_of_day_enum, trade_species_pokemon_id, turn_upside_down, spin, take_damage, critical_hits, updated_at)\n\tVALUES %s\nON CONFLICT (from_pokemon_id, to_pokemon_id, time_of_day_enum, gender_enum, trigger_enum)\n\tDO UPDATE SET\n\t\titem_id = EXCLUDED.item_id,\n\t\theld_item_id = EXCLUDED.held_item_id,\n\t\tknown_move_id = EXCLUDED.known_move_id,\n\t\tlocation_id = EXCLUDED.location_id,\n\t\tmin_level = EXCLUDED.min_level,\n\t\tmin_happiness = EXCLUDED.min_happiness,\n\t\tmin_beauty = EXCLUDED.min_beauty,\n\t\tmin_affection = EXCLUDED.min_affection,\n\t\tneeds_overworld_rain = EXCLUDED.needs_overworld_rain,\n\t\tparty_species_pokemon_id = EXCLUDED.party_species_pokemon_id,\n\t\tparty_type_id = EXCLUDED.party_type_id,\n\t\trelative_physical_stats = EXCLUDED.relative_physical_stats,\n\t\ttrade_species_pokemon_id = EXCLUDED.trade_species_pokemon_id,\n\t\tturn_upside_down = EXCLUDED.turn_upside_down,\n\t\tspin = EXCLUDED.spin,\n\t\ttake_damage = EXCLUDED.take_damage,\n\t\tcritical_hits = EXCLUDED.critical_hits,\n\t\tupdated_at = EXCLUDED.updated_at;",
		strings.Join(pokemonValues, ", "),
		strings.Join(pokemonTypeValues, ", "),
		strings.Join(pokemonAbilityValues, ", "),
		strings.Join(pokemonEggGroupValues, ", "),
		strings.Join(pokemonEvolutionValues, ", "),
	)

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()
	f.Close()

	moveChunks := helpers.Chunk(1, pokemonMoveValues)
	for i := range moveChunks {
		moveF := helpers.OpenFile(strings.Replace(config.OutputFile, ".sql", fmt.Sprintf("_moves_%d.sql", i), 1))
		moveSql := fmt.Sprintf("INSERT INTO pokemon_move (pokemon_id, move_id)\n\tVALUES %s\nON CONFLICT (pokemon_id, move_id)\n\tDO NOTHING;", strings.Join(moveChunks[i], ", "))
		moveO, err := moveF.WriteString(moveSql)
		if err != nil {
			log.Logger.Fatal(err)
		}

		moveF.Sync()
		moveF.Close()
		o = o + moveO
	}

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}

func parseEvolution(evolution pokeapi.Evolution, chain pokeapi.Evolution) []string {
	values := make([]string, 0)

	for _, details := range chain.EvolutionDetails {
		if details.Location == nil && (evolution.Species.Name != "feebas" || (evolution.Species.Name == "feebas" && details.Trigger.Name == "trade")) {
			gender := "'unknown'"
			if details.Gender == 1 {
				gender = "'male'"
			}
			if details.Gender == 2 {
				gender = "'female'"
			}

			item := "NULL"
			if details.Item != nil {
				item = fmt.Sprintf("(SELECT id from items WHERE slug='%s')", details.Item.Name)
			}

			heldItem := "NULL"
			if details.HeldItem != nil {
				heldItem = fmt.Sprintf("(SELECT id from items WHERE slug='%s')", details.HeldItem.Name)
			}

			knownMove := "NULL"
			if details.KnownMove != nil {
				knownMove = fmt.Sprintf("(SELECT id from moves WHERE slug='%s')", details.KnownMove.Name)
			}

			knownMoveType := "NULL"
			if details.KnownMoveType != nil {
				knownMoveType = fmt.Sprintf("(SELECT id from types WHERE slug='%s')", details.KnownMoveType.Name)
			}

			partySpecies := "NULL"
			if details.PartySpecies != nil {
				partySpecies = fmt.Sprintf("(SELECT id from pokemon WHERE slug='%s')", details.PartySpecies.Name)
			}

			partyType := "NULL"
			if details.PartyType != nil {
				partyType = fmt.Sprintf("(SELECT id from types WHERE slug='%s')", details.PartyType.Name)
			}

			tradeSpecies := "NULL"
			if details.TradeSpecies != nil {
				tradeSpecies = fmt.Sprintf("(SELECT id from pokemon WHERE slug='%s')", details.TradeSpecies.Name)
			}

			timeOfDay := "any"
			if details.TimeOfDay != "" {
				timeOfDay = details.TimeOfDay
			}

			toPokemonSlug := chain.Species.Name
			if toPokemonSlug == "meowstic" {
				toPokemonSlug = "meowstic-male"

				if details.Gender == 2 {
					toPokemonSlug = "meowstic-female"
				}
			}
			if toPokemonSlug == "toxtricity" {
				toPokemonSlug = "toxtricity-amped"
			}
			if toPokemonSlug == "urshifu" {
				toPokemonSlug = "urshifu-single-strike"
			}
			if toPokemonSlug == "lycanroc" {
				toPokemonSlug = "lycanroc-midday"

				if details.TimeOfDay == "night" {
					toPokemonSlug = "lycanroc-midnight"
				}
			}

			multipleSlugs := []string{toPokemonSlug}

			if toPokemonSlug == "toxtricity-amped" {
				multipleSlugs = append(multipleSlugs, "toxtricity-low-key")
			}
			if toPokemonSlug == "urshifu-single-strike" {
				multipleSlugs = append(multipleSlugs, "urshifu-rapid-strike")
			}

			for _, slug := range multipleSlugs {
				values = append(values, fmt.Sprintf(
					"(%s, %s, '%s', %s, %s, %s, %s, %s, %s, %d, %d, %d, %d, %t, %s, %s, %d, '%s', %s, %t, %t, %d, %d, now())",
					fmt.Sprintf("(SELECT id from pokemon WHERE slug='%s')", evolution.Species.Name),
					fmt.Sprintf("(SELECT id from pokemon WHERE slug='%s')", slug),
					strings.ReplaceAll(details.Trigger.Name, "_", "-"),
					item,
					gender,
					heldItem,
					knownMove,
					knownMoveType,
					"NULL",
					details.MinLevel,
					details.MinHappiness,
					details.MinBeauty,
					details.MinAffection,
					details.NeedsOverworldRain,
					partySpecies,
					partyType,
					details.RelativePhysicalStats,
					timeOfDay,
					tradeSpecies,
					details.TurnUpsideDown,
					false,
					0,
					0,
				))
			}

			for _, innerChain := range chain.EvolvesTo {
				evolutionValues := parseEvolution(chain, innerChain)
				values = append(values, evolutionValues...)
			}
		}
	}

	return values
}
