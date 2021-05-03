package main

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/pokeapi"
	"bekapod/pkmn-team-graphql/scripts/seed/helpers"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	start := time.Now()
	config := &helpers.Config{}
	arg.MustParse(config)

	client := pokeapi.New(pokeapi.PokeApiConfig{Host: config.Host, Prefix: config.Prefix})
	prisma := db.NewClient()
	if err := prisma.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := prisma.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	abilityList := client.GetResourceList("ability", 10000)
	results := abilityList.Results
	resultsLength := len(results)

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
				var effect *string
				if englishEffectEntry != nil {
					effect = &englishEffectEntry.ShortEffect
				}

				_, dbErr := prisma.Ability.UpsertOne(db.Ability.Slug.Equals(name)).
					Create(db.Ability.Slug.Set(name), db.Ability.Name.Set(englishName.Name), db.Ability.Effect.SetIfPresent(effect), db.Ability.UpdatedAt.Set(time.Now())).
					Update(db.Ability.Name.Set(englishName.Name), db.Ability.Effect.SetIfPresent(effect), db.Ability.UpdatedAt.Set(time.Now())).
					Exec(ctx)

				if dbErr != nil {
					log.Logger.Fatal(dbErr)
				}
			}
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Completed in %s\n", elapsed))
}
