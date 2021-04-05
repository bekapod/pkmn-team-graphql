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
	damageRelationValues := make([]string, 0)

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
				"('%s', '%s', now())",
				fullType.Name,
				helpers.EscapeSingleQuote(englishName.Name),
			))

			if len(fullType.DamageRelations.NoDamageTo) > 0 {
				for _, relation := range fullType.DamageRelations.NoDamageTo {
					damageRelationValues = append(damageRelationValues, fmt.Sprintf(
						"(%s, %s, 'no-damage-to', now())",
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", fullType.Name),
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", relation.Name),
					))
				}
			}

			if len(fullType.DamageRelations.HalfDamageTo) > 0 {
				for _, relation := range fullType.DamageRelations.HalfDamageTo {
					damageRelationValues = append(damageRelationValues, fmt.Sprintf(
						"(%s, %s, 'half-damage-to', now())",
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", fullType.Name),
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", relation.Name),
					))
				}
			}

			if len(fullType.DamageRelations.DoubleDamageTo) > 0 {
				for _, relation := range fullType.DamageRelations.DoubleDamageTo {
					damageRelationValues = append(damageRelationValues, fmt.Sprintf(
						"(%s, %s, 'double-damage-to', now())",
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", fullType.Name),
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", relation.Name),
					))
				}
			}

			if len(fullType.DamageRelations.NoDamageFrom) > 0 {
				for _, relation := range fullType.DamageRelations.NoDamageFrom {
					damageRelationValues = append(damageRelationValues, fmt.Sprintf(
						"(%s, %s, 'no-damage-from', now())",
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", fullType.Name),
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", relation.Name),
					))
				}
			}

			if len(fullType.DamageRelations.HalfDamageFrom) > 0 {
				for _, relation := range fullType.DamageRelations.HalfDamageFrom {
					damageRelationValues = append(damageRelationValues, fmt.Sprintf(
						"(%s, %s, 'half-damage-from', now())",
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", fullType.Name),
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", relation.Name),
					))
				}
			}

			if len(fullType.DamageRelations.DoubleDamageFrom) > 0 {
				for _, relation := range fullType.DamageRelations.DoubleDamageFrom {
					damageRelationValues = append(damageRelationValues, fmt.Sprintf(
						"(%s, %s, 'double-damage-from', now())",
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", fullType.Name),
						fmt.Sprintf("(SELECT id FROM types WHERE slug='%s')", relation.Name),
					))
				}
			}
		}(i)
	}

	wg.Wait()

	sql := fmt.Sprintf(
		"INSERT INTO types (slug, name, updated_at)\n\tVALUES %s\nON CONFLICT (slug)\n\tDO UPDATE SET\n\t\tname = EXCLUDED.name, updated_at = EXCLUDED.updated_at;\n\nINSERT INTO type_damage_relations (type_id, related_type_id, damage_relation_enum, updated_at)\n\tVALUES %s\nON CONFLICT (type_id, related_type_id, damage_relation_enum)\n\tDO UPDATE SET\n\t\tupdated_at = EXCLUDED.updated_at;",
		strings.Join(values, ", "),
		strings.Join(damageRelationValues, ", "),
	)

	o, err := f.WriteString(sql)
	if err != nil {
		log.Logger.Fatal(err)
	}
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
