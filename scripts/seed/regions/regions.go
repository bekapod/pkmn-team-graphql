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
		OutputFile: "seeds/regions.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)
	defer f.Close()

	regionList := client.GetResourceList("region")
	results := regionList.Results
	resultsLength := len(results)
	values := make([]string, 0)
	locationValues := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)
	sem := make(chan bool, 20)

	for i := 0; i < resultsLength; i++ {
		sem <- true
		go func(i int) {
			defer func() { <-sem }()
			defer wg.Done()
			fullRegion := client.GetRegion(results[i].Name)

			if fullRegion.MainGeneration.Name == "generation-viii" {
				englishName, err := pokeapi.GetEnglishName(fullRegion.Names, fullRegion.Name)
				if err != nil {
					log.Logger.Fatal(err)
				}

				values = append(values, fmt.Sprintf(
					"('%s', '%s', now())",
					fullRegion.Name,
					helpers.EscapeSingleQuote(englishName.Name),
				))

				locationsLength := len(fullRegion.Locations)
				var wg2 sync.WaitGroup
				wg2.Add(locationsLength)
				sem2 := make(chan bool, 20)

				for i := 0; i < locationsLength; i++ {
					sem2 <- true
					go func(i int) {
						defer func() { <-sem2 }()
						defer wg2.Done()
						urlParts := strings.Split(fullRegion.Locations[i].Url, "/")
						id := urlParts[len(urlParts)-2]
						fullLocation := client.GetLocation(id)
						locationEnglishName, err := pokeapi.GetEnglishName(fullLocation.Names, fullLocation.Name)
						if err != nil {
							log.Logger.Fatal(err)
						}

						locationValues = append(locationValues, fmt.Sprintf(
							"('%s', %s, '%s', now())",
							fullLocation.Name,
							fmt.Sprintf("(SELECT id from regions WHERE slug='%s')", fullLocation.Region.Name),
							helpers.EscapeSingleQuote(locationEnglishName.Name),
						))
					}(i)
				}
			}

		}(i)
	}

	wg.Wait()

	locationSql := ""

	if len(locationValues) != 0 {
		locationSql = fmt.Sprintf("\n\nINSERT INTO locations (slug, region_id, name, updated_at)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tregion_id = EXCLUDED.region_id,\n\t\tname = EXCLUDED.name,\n\t\tupdated_at = EXCLUDED.updated_at;", strings.Join(locationValues, ", "))
	}

	sql := fmt.Sprintf(
		"INSERT INTO regions (slug, name, updated_at)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tname = EXCLUDED.name,\n\t\t updated_at = EXCLUDED.updated_at;"+locationSql,
		strings.Join(values, ", "),
	)

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
