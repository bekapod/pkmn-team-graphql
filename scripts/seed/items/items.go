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

	itemList := client.GetResourceList("item", 10000)
	results := itemList.Results
	resultsLength := len(results)

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
			attributes := make([]db.ItemAttribute, 0)

			for _, attribute := range fullItem.Attributes {
				attributes = append(attributes, db.ItemAttribute(attribute.Name))
			}

			var effect *string
			if englishEffectEntry != nil {
				effect = &englishEffectEntry.ShortEffect
			}

			var flingEffect *string
			if fullItem.FlingEffect != nil {
				effect = &fullItem.FlingEffect.Name
			}

			_, dbErr := prisma.Item.UpsertOne(db.Item.Slug.Equals(fullItem.Name)).
				Create(
					db.Item.Slug.Set(fullItem.Name),
					db.Item.Name.Set(englishName.Name),
					db.Item.Category.Set(db.ItemCategory(fullItem.Category.Name)),
					db.Item.Attributes.Set(attributes),
					db.Item.Cost.SetIfPresent(fullItem.Cost),
					db.Item.FlingPower.SetIfPresent(fullItem.FlingPower),
					db.Item.Effect.SetIfPresent(effect),
					db.Item.FlingEffect.SetIfPresent(flingEffect),
					db.Item.Sprite.SetIfPresent(fullItem.Sprites.Default),
					db.Item.UpdatedAt.Set(time.Now())).
				Update(
					db.Item.Name.Set(englishName.Name),
					db.Item.Category.Set(db.ItemCategory(fullItem.Category.Name)),
					db.Item.Attributes.Set(attributes),
					db.Item.Cost.SetIfPresent(fullItem.Cost),
					db.Item.FlingPower.SetIfPresent(fullItem.FlingPower),
					db.Item.Effect.SetIfPresent(effect),
					db.Item.FlingEffect.SetIfPresent(flingEffect),
					db.Item.Sprite.SetIfPresent(fullItem.Sprites.Default),
					db.Item.UpdatedAt.Set(time.Now())).
				Exec(ctx)

			if dbErr != nil {
				log.Logger.Fatal(dbErr)
			}
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Completed in %s\n", elapsed))
}
