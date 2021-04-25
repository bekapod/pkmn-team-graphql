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
		OutputFile: "seeds/items.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)
	defer f.Close()

	itemList := client.GetResourceList("item")
	results := itemList.Results
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
			fullItem := client.GetItem(results[i].Name)

			englishName, err := pokeapi.GetEnglishName(fullItem.Names, fullItem.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			englishEffectEntry, err := pokeapi.GetEnglishEffectEntry(fullItem.EffectEntries, fullItem.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			attributes := make([]string, 0)

			for _, attribute := range fullItem.Attributes {
				attributes = append(attributes, strings.ReplaceAll(attribute.Name, "_", "-"))
			}

			flingEffect := "NULL"
			if fullItem.FlingEffect.Name != "" {
				flingEffect = fmt.Sprintf("'%s'", fullItem.FlingEffect.Name)
			}

			values = append(values, fmt.Sprintf(
				"('%s', '%s', %d, %d, %s, '%s', '%s', '%s', '%s', now())",
				fullItem.Name,
				helpers.EscapeSingleQuote(englishName.Name),
				fullItem.Cost,
				fullItem.FlingPower,
				flingEffect,
				helpers.EscapeSingleQuote(englishEffectEntry.ShortEffect),
				fullItem.Sprites.Default,
				strings.ReplaceAll(fullItem.Category.Name, "_", "-"),
				fmt.Sprintf("{%s}", strings.Join(attributes, ",")),
			))

		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf("INSERT INTO items (slug, name, cost, fling_power, fling_effect, effect, sprite, category_enum, attribute_enums, updated_at)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tname = EXCLUDED.name,\n\t\tcost = EXCLUDED.cost,\n\t\tfling_power = EXCLUDED.fling_power,\n\t\tfling_effect = EXCLUDED.fling_effect,\n\t\teffect = EXCLUDED.effect,\n\t\tsprite = EXCLUDED.sprite,\n\t\tcategory_enum = EXCLUDED.category_enum,\n\t\tattribute_enums = EXCLUDED.attribute_enums,\n\t\t updated_at = EXCLUDED.updated_at;", strings.Join(values, ", "))

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
