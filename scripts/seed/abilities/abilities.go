package main

import (
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/scripts"
	"bekapod/pkmn-team-graphql/scripts/pokeapi"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
)

func getAllAbilities() pokeapi.ResourcePointerList {
	abilityList := pokeapi.ResourcePointerList{}
	scripts.GetResource("ability?limit=1000", &abilityList)
	return abilityList
}

func getAbility(id string) pokeapi.RawAbility {
	fullAbility := pokeapi.RawAbility{}
	scripts.GetResource(fmt.Sprintf("ability/%s", id), &fullAbility)
	return fullAbility
}

func generateAbilitiesSeed() string {
	abilityList := getAllAbilities()
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
			fullAbility := getAbility(id)
			if fullAbility.IsMainSeries {
				englishName, err := pokeapi.GetEnglishName(fullAbility.Names, fullAbility.Name)
				scripts.Check(err)
				englishEffectEntry, _ := pokeapi.GetEnglishEffectEntry(fullAbility.EffectEntries, fullAbility.Name)

				values = append(values, fmt.Sprintf(
					"('%s', '%s', '%s')",
					scripts.EscapeSingleQuote(fullAbility.Name),
					scripts.EscapeSingleQuote(englishName.Name),
					scripts.EscapeSingleQuote(englishEffectEntry.ShortEffect),
				))
			}
		}(i)
	}

	wg.Wait()

	return fmt.Sprintf("INSERT INTO abilities (slug, name, effect)\n\tVALUES %s;", strings.Join(values, ", "))
}

func main() {
	start := time.Now()
	scripts.Args = &scripts.SeedArgs{
		OutputFile: "seeds/abilities.sql",
	}
	arg.MustParse(scripts.Args)

	f := scripts.OpenFile(scripts.Args.OutputFile)
	defer f.Close()
	sql := generateAbilitiesSeed()

	o, err := f.WriteString(sql)
	scripts.Check(err)
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
