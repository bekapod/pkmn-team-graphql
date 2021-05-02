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

	typeList := client.GetResourceList("type", 10000)
	results := typeList.Results
	resultsLength := len(results)
	types := make([]pokeapi.Type, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			urlParts := strings.Split(results[i].Url, "/")
			id := urlParts[len(urlParts)-2]
			fullType := client.GetType(id)
			types = append(types, fullType)
			englishName, err := pokeapi.GetEnglishName(fullType.Names, fullType.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}

			_, dbErr := prisma.Type.UpsertOne(db.Type.Slug.Equals(fullType.Name)).
				Create(db.Type.Slug.Set(fullType.Name), db.Type.Name.Set(englishName.Name), db.Type.UpdatedAt.Set(time.Now())).
				Update(db.Type.Name.Set(englishName.Name), db.Type.UpdatedAt.Set(time.Now())).
				Exec(ctx)

			if dbErr != nil {
				log.Logger.Fatal(dbErr)
			}
		}(i)
	}

	wg.Wait()

	for _, fullType := range types {
		if len(fullType.DamageRelations.NoDamageTo) > 0 {
			for _, relation := range fullType.DamageRelations.NoDamageTo {
				_, err := prisma.TypeDamageRelation.CreateOne(
					db.TypeDamageRelation.TypeA.Link(db.Type.Slug.Equals(fullType.Name)),
					db.TypeDamageRelation.TypeB.Link(db.Type.Slug.Equals(relation.Name)),
					db.TypeDamageRelation.DamageRelation.Set("NO_DAMAGE_TO"),
					db.TypeDamageRelation.UpdatedAt.Set(time.Now()),
				).Exec(ctx)

				if err != nil {
					log.Logger.Fatal(err)
				}
			}
		}

		if len(fullType.DamageRelations.HalfDamageTo) > 0 {
			for _, relation := range fullType.DamageRelations.HalfDamageTo {
				_, err := prisma.TypeDamageRelation.CreateOne(
					db.TypeDamageRelation.TypeA.Link(db.Type.Slug.Equals(fullType.Name)),
					db.TypeDamageRelation.TypeB.Link(db.Type.Slug.Equals(relation.Name)),
					db.TypeDamageRelation.DamageRelation.Set("HALF_DAMAGE_TO"),
					db.TypeDamageRelation.UpdatedAt.Set(time.Now()),
				).Exec(ctx)

				if err != nil {
					log.Logger.Fatal(err)
				}
			}
		}

		if len(fullType.DamageRelations.DoubleDamageTo) > 0 {
			for _, relation := range fullType.DamageRelations.DoubleDamageTo {
				_, err := prisma.TypeDamageRelation.CreateOne(
					db.TypeDamageRelation.TypeA.Link(db.Type.Slug.Equals(fullType.Name)),
					db.TypeDamageRelation.TypeB.Link(db.Type.Slug.Equals(relation.Name)),
					db.TypeDamageRelation.DamageRelation.Set("DOUBLE_DAMAGE_TO"),
					db.TypeDamageRelation.UpdatedAt.Set(time.Now()),
				).Exec(ctx)

				if err != nil {
					log.Logger.Fatal(err)
				}
			}
		}

		if len(fullType.DamageRelations.NoDamageFrom) > 0 {
			for _, relation := range fullType.DamageRelations.NoDamageFrom {
				_, err := prisma.TypeDamageRelation.CreateOne(
					db.TypeDamageRelation.TypeA.Link(db.Type.Slug.Equals(fullType.Name)),
					db.TypeDamageRelation.TypeB.Link(db.Type.Slug.Equals(relation.Name)),
					db.TypeDamageRelation.DamageRelation.Set("NO_DAMAGE_FROM"),
					db.TypeDamageRelation.UpdatedAt.Set(time.Now()),
				).Exec(ctx)

				if err != nil {
					log.Logger.Fatal(err)
				}
			}
		}

		if len(fullType.DamageRelations.HalfDamageFrom) > 0 {
			for _, relation := range fullType.DamageRelations.HalfDamageFrom {
				_, err := prisma.TypeDamageRelation.CreateOne(
					db.TypeDamageRelation.TypeA.Link(db.Type.Slug.Equals(fullType.Name)),
					db.TypeDamageRelation.TypeB.Link(db.Type.Slug.Equals(relation.Name)),
					db.TypeDamageRelation.DamageRelation.Set("HALF_DAMAGE_FROM"),
					db.TypeDamageRelation.UpdatedAt.Set(time.Now()),
				).Exec(ctx)

				if err != nil {
					log.Logger.Fatal(err)
				}
			}
		}

		if len(fullType.DamageRelations.DoubleDamageFrom) > 0 {
			for _, relation := range fullType.DamageRelations.DoubleDamageFrom {
				_, err := prisma.TypeDamageRelation.CreateOne(
					db.TypeDamageRelation.TypeA.Link(db.Type.Slug.Equals(fullType.Name)),
					db.TypeDamageRelation.TypeB.Link(db.Type.Slug.Equals(relation.Name)),
					db.TypeDamageRelation.DamageRelation.Set("DOUBLE_DAMAGE_FROM"),
					db.TypeDamageRelation.UpdatedAt.Set(time.Now()),
				).Exec(ctx)

				if err != nil {
					log.Logger.Fatal(err)
				}
			}
		}
	}

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Completed in %s\n", elapsed))
}
