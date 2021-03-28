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

			englishName, err := pokeapi.GetEnglishName(fullPokemonSpecies.Names, fullPokemonSpecies.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			englishFlavourText, _ := pokeapi.GetEnglishFlavourTextEntry(fullPokemonSpecies.FlavourTextEntries, fullPokemonSpecies.Name)

			for i := range varieties {
				pokemon := varieties[i]
				pokemonValues = append(pokemonValues, fmt.Sprintf(
					"(%d, '%s', '%s', '%s', %d, %d, %d, %d, %d, %d, %t, %t, %t, '%s')",
					fullPokemonSpecies.ID,
					pokemon.Name,
					helpers.EscapeSingleQuote(englishName.Name),
					pokemon.Sprites.FrontDefault,
					pokeapi.GetPokemonStat(pokemon, "hp"),
					pokeapi.GetPokemonStat(pokemon, "attack"),
					pokeapi.GetPokemonStat(pokemon, "defense"),
					pokeapi.GetPokemonStat(pokemon, "special-attack"),
					pokeapi.GetPokemonStat(pokemon, "special-defense"),
					pokeapi.GetPokemonStat(pokemon, "speed"),
					fullPokemonSpecies.IsBaby,
					fullPokemonSpecies.IsLegendary,
					fullPokemonSpecies.IsMythical,
					helpers.EscapeSingleQuote(englishFlavourText.FlavourText),
				))

				for i := range pokemon.Types {
					pokemonTypeValues = append(pokemonTypeValues, fmt.Sprintf(
						"(%s, %s)",
						fmt.Sprintf("(SELECT id from pokemon WHERE slug='%s')", pokemon.Name),
						fmt.Sprintf("(SELECT id from types WHERE slug='%s')", pokemon.Types[i].Type.Name),
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
			}
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf(
		"INSERT INTO pokemon (pokedex_id, slug, name, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description)\n\tVALUES %s;\n\n"+
			"INSERT INTO pokemon_type (pokemon_id, type_id)\n\tVALUES %s;\n\n"+
			"INSERT INTO pokemon_ability (pokemon_id, ability_id)\n\tVALUES %s;\n\n",
		strings.Join(pokemonValues, ", "),
		strings.Join(pokemonTypeValues, ", "),
		strings.Join(pokemonAbilityValues, ", "),
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
		moveSql := fmt.Sprintf("INSERT INTO pokemon_move (pokemon_id, move_id)\n\tVALUES %s;", strings.Join(moveChunks[i], ", "))
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
