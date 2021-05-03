package main

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/pokeapi"
	"bekapod/pkmn-team-graphql/scripts/seed/helpers"
	"context"
	"fmt"
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

	moveList := client.GetResourceList("move", 10000)
	results := moveList.Results
	resultsLength := len(results)

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
				englishName, err := pokeapi.GetEnglishName(fullMove.Names, fullMove.Name)
				if err != nil {
					log.Logger.Fatal(err)
				}
				englishEffectEntry, _ := pokeapi.GetEnglishEffectEntry(fullMove.EffectEntries, fullMove.Name)
				var effect *string
				if englishEffectEntry != nil {
					effect = &englishEffectEntry.ShortEffect
				}

				_, dbErr := prisma.Move.UpsertOne(db.Move.Slug.Equals(fullMove.Name)).
					Create(
						db.Move.Slug.Set(fullMove.Name),
						db.Move.Name.Set(englishName.Name),
						db.Move.DamageClass.Set(db.DamageClass(fullMove.DamageClass.Name)),
						db.Move.Target.Set(db.MoveTarget(fullMove.Target.Name)),
						db.Move.Type.Link(db.Type.Slug.Equals(fullMove.Type.Name)),
						db.Move.Effect.SetIfPresent(effect),
						db.Move.EffectChance.SetIfPresent(fullMove.EffectChance),
						db.Move.Accuracy.SetIfPresent(fullMove.Accuracy),
						db.Move.Pp.SetIfPresent(fullMove.PP),
						db.Move.Power.SetIfPresent(fullMove.Power),
						db.Move.UpdatedAt.Set(time.Now())).
					Update(
						db.Move.Name.Set(englishName.Name),
						db.Move.DamageClass.Set(db.DamageClass(fullMove.DamageClass.Name)),
						db.Move.Target.Set(db.MoveTarget(fullMove.Target.Name)),
						db.Move.Type.Link(db.Type.Slug.Equals(fullMove.Type.Name)),
						db.Move.Effect.SetIfPresent(effect),
						db.Move.EffectChance.SetIfPresent(fullMove.EffectChance),
						db.Move.Accuracy.SetIfPresent(fullMove.Accuracy),
						db.Move.Pp.SetIfPresent(fullMove.PP),
						db.Move.Power.SetIfPresent(fullMove.Power),
						db.Move.UpdatedAt.Set(time.Now())).
					Exec(ctx)

				if dbErr != nil {
					log.Logger.WithField("move", fullMove.Name).Fatal(dbErr)
				}
			}
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Completed in %s\n", elapsed))
}
