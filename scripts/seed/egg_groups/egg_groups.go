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

	eggGroupList := client.GetResourceList("egg-group", 10000)
	results := eggGroupList.Results
	resultsLength := len(results)

	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			urlParts := strings.Split(results[i].Url, "/")
			id := urlParts[len(urlParts)-2]
			fullEggGroup := client.GetEggGroup(id)

			englishName, err := pokeapi.GetEnglishName(fullEggGroup.Names, fullEggGroup.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}

			_, dbErr := prisma.EggGroup.UpsertOne(db.EggGroup.Slug.Equals(fullEggGroup.Name)).
				Create(db.EggGroup.Slug.Set(fullEggGroup.Name), db.EggGroup.Name.Set(englishName.Name), db.EggGroup.UpdatedAt.Set(time.Now())).
				Update(db.EggGroup.Name.Set(englishName.Name), db.EggGroup.UpdatedAt.Set(time.Now())).
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
