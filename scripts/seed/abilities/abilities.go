// +build tools

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
		OutputFile: "seeds/abilities.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)
	defer f.Close()

	abilityList := client.GetResourceList("ability")
	results := abilityList.Results
	resultsLength := len(results)
	values := make([]string, 0)

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
			fullAbility := client.GetAbility(id)
			name := fullAbility.Name

			if len(fullAbility.Pokemon) == 1 && fullAbility.Pokemon[0].Pokemon.Name == "calyrex-shadow-rider" {
				name = "as-one-shadow-rider"
			}

			if len(fullAbility.Pokemon) == 1 && fullAbility.Pokemon[0].Pokemon.Name == "calyrex-ice-rider" {
				name = "as-one-ice-rider"
			}

			if fullAbility.IsMainSeries {
				englishName, err := pokeapi.GetEnglishName(fullAbility.Names, fullAbility.Name)
				if err != nil {
					log.Logger.Fatal(err)
				}
				englishEffectEntry, _ := pokeapi.GetEnglishEffectEntry(fullAbility.EffectEntries, fullAbility.Name)

				values = append(values, fmt.Sprintf(
					"('%s', '%s', '%s')",
					name,
					helpers.EscapeSingleQuote(englishName.Name),
					helpers.EscapeSingleQuote(englishEffectEntry.ShortEffect),
				))
			}
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf("INSERT INTO abilities (slug, name, effect)\n\tVALUES %s;", strings.Join(values, ", "))

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
