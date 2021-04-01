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
		OutputFile: "seeds/types.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)
	defer f.Close()

	typeList := client.GetResourceList("type")
	results := typeList.Results
	resultsLength := len(results)
	values := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			urlParts := strings.Split(results[i].Url, "/")
			id := urlParts[len(urlParts)-2]
			fullType := client.GetType(id)

			englishName, err := pokeapi.GetEnglishName(fullType.Names, fullType.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}

			values = append(values, fmt.Sprintf(
				"('%s', '%s')",
				fullType.Name,
				helpers.EscapeSingleQuote(englishName.Name),
			))
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf("INSERT INTO types (slug, name)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tname = EXCLUDED.name;", strings.Join(values, ", "))

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
