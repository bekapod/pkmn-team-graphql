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
		OutputFile: "seeds/stats.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)
	defer f.Close()

	statList := client.GetResourceList("stat")
	results := statList.Results
	resultsLength := len(results)
	values := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			fullStat := client.GetStat(results[i].Name)
			englishName, err := pokeapi.GetEnglishName(fullStat.Names, fullStat.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}

			values = append(values, fmt.Sprintf(
				"('%s', '%s')",
				fullStat.Name,
				helpers.EscapeSingleQuote(englishName.Name),
			))
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf("INSERT INTO stats (slug, name)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tname = EXCLUDED.name;", strings.Join(values, ", "))

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
