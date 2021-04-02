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

func main() {
	start := time.Now()
	config := &helpers.Config{
		OutputFile: "seeds/pokemon.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)

	pokemonSpeciesList := client.GetResourceList("pokemon-species")
	results := pokemonSpeciesList.Results
	resultsLength := len(results)
	pokemonValues := make([]string, 0)
	pokemonTypeValues := make([]string, 0)
	pokemonAbilityValues := make([]string, 0)
	pokemonMoveValues := make([]string, 0)
	pokemonEggGroupValues := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)
	sem := make(chan bool, 20)

	for i := 0; i < resultsLength; i++ {
		sem <- true
		go func(i int) {
			defer func() { <-sem }()
			defer wg.Done()
			urlParts := strings.Split(results[i].Url, "/")
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
					"(%d, '%s', '%s', '%s', %d, %d, %d, %d, %d, %d, %t, %t, %t, '%s', '%s', '%s', %s, %t, '%s', %d, %d)",
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
						"(%s, %s, %d)",
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
						"(%s, %s)",
						fmt.Sprintf("(SELECT id FROM pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id FROM abilities WHERE slug='%s')", abilityName),
					))
				}

				for i := range pokemon.Moves {
					pokemonMoveValues = append(pokemonMoveValues, fmt.Sprintf(
						"(%s, %s)",
						fmt.Sprintf("(SELECT id FROM pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id FROM moves WHERE slug='%s')", pokemon.Moves[i].Move.Name),
					))
				}

				for i := range fullPokemonSpecies.EggGroups {
					pokemonEggGroupValues = append(pokemonEggGroupValues, fmt.Sprintf(
						"(%s, %s)",
						fmt.Sprintf("(SELECT id FROM pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id FROM egg_groups WHERE slug='%s')", fullPokemonSpecies.EggGroups[i].Name),
					))
				}
			}
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf(
		"INSERT INTO pokemon (pokedex_id, slug, name, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, color_enum, shape_enum, habitat_enum, is_default_variant, genus, height, weight)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tpokedex_id = EXCLUDED.pokedex_id,\n\t\tname = EXCLUDED.name,\n\t\tsprite = EXCLUDED.sprite,\n\t\thp = EXCLUDED.hp,\n\t\tattack = EXCLUDED.attack,\n\t\tdefense = EXCLUDED.defense,\n\t\tspecial_attack = EXCLUDED.special_attack,\n\t\tspecial_defense = EXCLUDED.special_defense,\n\t\tspeed = EXCLUDED.speed,\n\t\tis_baby = EXCLUDED.is_baby,\n\t\tis_legendary = EXCLUDED.is_legendary,\n\t\tis_mythical = EXCLUDED.is_mythical,\n\t\tdescription = EXCLUDED.description,\n\t\tcolor_enum = EXCLUDED.color_enum,\n\t\tshape_enum = EXCLUDED.shape_enum,\n\t\thabitat_enum = EXCLUDED.habitat_enum,\n\t\tis_default_variant = EXCLUDED.is_default_variant,\n\t\tgenus = EXCLUDED.genus,\n\t\theight = EXCLUDED.height,\n\t\tweight = EXCLUDED.weight;\n\n"+
			"INSERT INTO pokemon_type (pokemon_id, type_id, slot)\n\tVALUES %s\nON CONFLICT (pokemon_id, type_id)\n\tDO UPDATE SET\n\t\tslot = EXCLUDED.slot;\n\n"+
			"INSERT INTO pokemon_ability (pokemon_id, ability_id)\n\tVALUES %s\nON CONFLICT (pokemon_id, ability_id)\n\tDO NOTHING;\n\n"+
			"INSERT INTO pokemon_egg_group (pokemon_id, egg_group_id)\n\tVALUES %s\nON CONFLICT (pokemon_id, egg_group_id)\n\tDO NOTHING;",
		strings.Join(pokemonValues, ", "),
		strings.Join(pokemonTypeValues, ", "),
		strings.Join(pokemonAbilityValues, ", "),
		strings.Join(pokemonEggGroupValues, ", "),
	)

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()
	f.Close()

	moveChunks := helpers.Chunk(10, pokemonMoveValues)
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
