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

func getListOfMoveTargets(client *pokeapi.PokeApiClient) map[string]string {
	moveTargetList := client.GetResourceList("move-target")
	results := moveTargetList.Results
	resultsLength := len(results)
	targets := make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			target := client.GetMoveTarget(results[i].Name)
			englishName, err := pokeapi.GetEnglishName(target.Names, target.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			targets[target.Name] = englishName.Name
		}(i)
	}

	wg.Wait()
	return targets
}

func main() {
	start := time.Now()
	config := &helpers.Config{
		OutputFile: "seeds/moves.sql",
	}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})

	f := helpers.OpenFile(config.OutputFile)
	defer f.Close()

	moveList := client.GetResourceList("move")
	targets := getListOfMoveTargets(client)
	results := moveList.Results
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
			fullMove := client.GetMove(results[i].Name)
			if fullMove.DamageClass.Name != "" {
				target := targets[fullMove.Target.Name]

				englishName, err := pokeapi.GetEnglishName(fullMove.Names, fullMove.Name)
				if err != nil {
					log.Logger.Fatal(err)
				}
				englishEffectEntry, _ := pokeapi.GetEnglishEffectEntry(fullMove.EffectEntries, fullMove.Name)

				values = append(values, fmt.Sprintf(
					"('%s', '%s', %d, %d, %d, '%s', '%s', %d, '%s', %s, now())",
					fullMove.Name,
					helpers.EscapeSingleQuote(englishName.Name),
					fullMove.Accuracy,
					fullMove.PP,
					fullMove.Power,
					helpers.EscapeSingleQuote(fullMove.DamageClass.Name),
					helpers.EscapeSingleQuote(englishEffectEntry.ShortEffect),
					fullMove.EffectChance,
					helpers.EscapeSingleQuote(target),
					fmt.Sprintf("(SELECT id from types WHERE slug='%s')", fullMove.Type.Name),
				))
			}
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf("INSERT INTO moves (slug, name, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id, updated_at)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tname = EXCLUDED.name,\n\t\taccuracy = EXCLUDED.accuracy,\n\t\tpp = EXCLUDED.pp,\n\t\tpower = EXCLUDED.power,\n\t\tdamage_class_enum = EXCLUDED.damage_class_enum,\n\t\teffect = EXCLUDED.effect,\n\t\teffect_chance = EXCLUDED.effect_chance,\n\t\ttarget = EXCLUDED.target,\n\t\ttype_id = EXCLUDED.type_id, updated_at = EXCLUDED.updated_at;", strings.Join(values, ", "))

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
