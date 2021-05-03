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

	generation := client.GetGeneration("generation-viii")
	pokemonSpeciesList := make([]string, 0)
	for _, versionGroupResource := range generation.VersionGroups {
		versionGroup := client.GetVersionGroup(versionGroupResource.Name)

		for _, pokedexResource := range versionGroup.Pokedexes {
			pokedex := client.GetPokedex(pokedexResource.Name)

			for _, pokemonEntry := range pokedex.PokemonEntries {
				pokemonSpeciesList = append(pokemonSpeciesList, pokemonEntry.PokemonSpecies.Url)
			}
		}
	}

	pokemonSpeciesList = append(pokemonSpeciesList, []string{
		"http://localhost/api/v2/pokemon-species/1/",
		"http://localhost/api/v2/pokemon-species/2/",
		"http://localhost/api/v2/pokemon-species/3/",
		"http://localhost/api/v2/pokemon-species/7/",
		"http://localhost/api/v2/pokemon-species/8/",
		"http://localhost/api/v2/pokemon-species/9/",
		"http://localhost/api/v2/pokemon-species/150/",
		"http://localhost/api/v2/pokemon-species/151/",
		"http://localhost/api/v2/pokemon-species/240/",
		"http://localhost/api/v2/pokemon-species/243/",
		"http://localhost/api/v2/pokemon-species/244/",
		"http://localhost/api/v2/pokemon-species/245/",
		"http://localhost/api/v2/pokemon-species/249/",
		"http://localhost/api/v2/pokemon-species/250/",
		"http://localhost/api/v2/pokemon-species/252/",
		"http://localhost/api/v2/pokemon-species/253/",
		"http://localhost/api/v2/pokemon-species/254/",
		"http://localhost/api/v2/pokemon-species/255/",
		"http://localhost/api/v2/pokemon-species/256/",
		"http://localhost/api/v2/pokemon-species/257/",
		"http://localhost/api/v2/pokemon-species/258/",
		"http://localhost/api/v2/pokemon-species/259/",
		"http://localhost/api/v2/pokemon-species/260/",
		"http://localhost/api/v2/pokemon-species/276/",
		"http://localhost/api/v2/pokemon-species/380/",
		"http://localhost/api/v2/pokemon-species/381/",
		"http://localhost/api/v2/pokemon-species/382/",
		"http://localhost/api/v2/pokemon-species/383/",
		"http://localhost/api/v2/pokemon-species/384/",
		"http://localhost/api/v2/pokemon-species/385/",
		"http://localhost/api/v2/pokemon-species/480/",
		"http://localhost/api/v2/pokemon-species/481/",
		"http://localhost/api/v2/pokemon-species/482/",
		"http://localhost/api/v2/pokemon-species/483/",
		"http://localhost/api/v2/pokemon-species/484/",
		"http://localhost/api/v2/pokemon-species/485/",
		"http://localhost/api/v2/pokemon-species/486/",
		"http://localhost/api/v2/pokemon-species/487/",
		"http://localhost/api/v2/pokemon-species/488/",
		"http://localhost/api/v2/pokemon-species/641/",
		"http://localhost/api/v2/pokemon-species/642/",
		"http://localhost/api/v2/pokemon-species/643/",
		"http://localhost/api/v2/pokemon-species/644/",
		"http://localhost/api/v2/pokemon-species/645/",
		"http://localhost/api/v2/pokemon-species/646/",
		"http://localhost/api/v2/pokemon-species/647/",
		"http://localhost/api/v2/pokemon-species/716/",
		"http://localhost/api/v2/pokemon-species/717/",
		"http://localhost/api/v2/pokemon-species/718/",
		"http://localhost/api/v2/pokemon-species/722/",
		"http://localhost/api/v2/pokemon-species/723/",
		"http://localhost/api/v2/pokemon-species/724/",
		"http://localhost/api/v2/pokemon-species/725/",
		"http://localhost/api/v2/pokemon-species/726/",
		"http://localhost/api/v2/pokemon-species/727/",
		"http://localhost/api/v2/pokemon-species/728/",
		"http://localhost/api/v2/pokemon-species/729/",
		"http://localhost/api/v2/pokemon-species/730/",
		"http://localhost/api/v2/pokemon-species/785/",
		"http://localhost/api/v2/pokemon-species/786/",
		"http://localhost/api/v2/pokemon-species/787/",
		"http://localhost/api/v2/pokemon-species/788/",
		"http://localhost/api/v2/pokemon-species/789/",
		"http://localhost/api/v2/pokemon-species/790/",
		"http://localhost/api/v2/pokemon-species/791/",
		"http://localhost/api/v2/pokemon-species/792/",
		"http://localhost/api/v2/pokemon-species/793/",
		"http://localhost/api/v2/pokemon-species/794/",
		"http://localhost/api/v2/pokemon-species/795/",
		"http://localhost/api/v2/pokemon-species/796/",
		"http://localhost/api/v2/pokemon-species/797/",
		"http://localhost/api/v2/pokemon-species/798/",
		"http://localhost/api/v2/pokemon-species/799/",
		"http://localhost/api/v2/pokemon-species/800/",
		"http://localhost/api/v2/pokemon-species/803/",
		"http://localhost/api/v2/pokemon-species/804/",
		"http://localhost/api/v2/pokemon-species/805/",
		"http://localhost/api/v2/pokemon-species/806/",
		"http://localhost/api/v2/pokemon-species/807/",
		"http://localhost/api/v2/pokemon-species/808/",
		"http://localhost/api/v2/pokemon-species/809/",
	}...)

	pokemonSpeciesList = unique(pokemonSpeciesList)
	resultsLength := len(pokemonSpeciesList)
	pokemonEvolutionChainList := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)
	sem := make(chan bool, 20)

	for i := 0; i < resultsLength; i++ {
		sem <- true
		go func(i int) {
			defer func() { <-sem }()
			defer wg.Done()
			urlParts := strings.Split(pokemonSpeciesList[i], "/")
			id := urlParts[len(urlParts)-2]
			fullPokemonSpecies := client.GetPokemonSpecies(id)
			eggGroups := make([]db.EggGroupWhereParam, 0)
			varieties := client.GetPokemonVarietiesForSpecies(fullPokemonSpecies)

			pokedexId, err := pokeapi.GetPokedexId(fullPokemonSpecies, "national")
			if err != nil {
				log.Logger.Fatal(err)
			}
			englishName, err := pokeapi.GetEnglishName(fullPokemonSpecies.Names, fullPokemonSpecies.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}
			englishFlavourText, _ := pokeapi.GetEnglishFlavourTextEntry(fullPokemonSpecies.FlavourTextEntries, fullPokemonSpecies.Name)
			var description *string
			if englishFlavourText != nil {
				description = &englishFlavourText.FlavourText
			}
			englishGenus, err := pokeapi.GetEnglishGenus(fullPokemonSpecies.Genera, fullPokemonSpecies.Name)
			if err != nil {
				log.Logger.Fatal(err)
			}

			for _, eggGroup := range fullPokemonSpecies.EggGroups {
				eggGroups = append(eggGroups, db.EggGroup.Slug.Equals(eggGroup.Name))
			}

			for i := range varieties {
				pokemon := varieties[i]
				hp, _ := pokeapi.GetPokemonStat(pokemon, "hp")
				attack, _ := pokeapi.GetPokemonStat(pokemon, "attack")
				defense, _ := pokeapi.GetPokemonStat(pokemon, "defense")
				specialAttack, _ := pokeapi.GetPokemonStat(pokemon, "special-attack")
				specialDefense, _ := pokeapi.GetPokemonStat(pokemon, "special-defense")
				speed, _ := pokeapi.GetPokemonStat(pokemon, "speed")
				var habitat *db.Habitat

				if fullPokemonSpecies.Habitat != nil {
					habitat = (*db.Habitat)(&fullPokemonSpecies.Habitat.Name)
				}

				createdPokemon, dbErr := prisma.Pokemon.UpsertOne(db.Pokemon.Slug.Equals(pokemon.Name)).
					Create(
						db.Pokemon.PokedexID.Set(pokedexId),
						db.Pokemon.Slug.Set(pokemon.Name),
						db.Pokemon.Name.Set(englishName.Name),
						db.Pokemon.Hp.Set(hp),
						db.Pokemon.Attack.Set(attack),
						db.Pokemon.Defense.Set(defense),
						db.Pokemon.SpecialAttack.Set(specialAttack),
						db.Pokemon.SpecialDefense.Set(specialDefense),
						db.Pokemon.Speed.Set(speed),
						db.Pokemon.IsBaby.Set(fullPokemonSpecies.IsBaby),
						db.Pokemon.IsLegendary.Set(fullPokemonSpecies.IsLegendary),
						db.Pokemon.IsMythical.Set(fullPokemonSpecies.IsMythical),
						db.Pokemon.Color.Set(db.Color(fullPokemonSpecies.Color.Name)),
						db.Pokemon.Shape.Set(db.Shape(fullPokemonSpecies.Shape.Name)),
						db.Pokemon.IsDefaultVariant.Set(pokemon.IsDefault),
						db.Pokemon.Genus.Set(englishGenus.Genus),
						db.Pokemon.Height.Set(pokemon.Height),
						db.Pokemon.Weight.Set(pokemon.Weight),
						db.Pokemon.Habitat.SetIfPresent(habitat),
						db.Pokemon.Description.SetIfPresent(description),
						db.Pokemon.Sprite.SetIfPresent(&pokemon.Sprites.FrontDefault),
						db.Pokemon.EggGroups.Link(eggGroups...),
						db.Pokemon.UpdatedAt.Set(time.Now())).
					Update(
						db.Pokemon.PokedexID.Set(pokedexId),
						db.Pokemon.Name.Set(englishName.Name),
						db.Pokemon.Hp.Set(hp),
						db.Pokemon.Attack.Set(attack),
						db.Pokemon.Defense.Set(defense),
						db.Pokemon.SpecialAttack.Set(specialAttack),
						db.Pokemon.SpecialDefense.Set(specialDefense),
						db.Pokemon.Speed.Set(speed),
						db.Pokemon.IsBaby.Set(fullPokemonSpecies.IsBaby),
						db.Pokemon.IsLegendary.Set(fullPokemonSpecies.IsLegendary),
						db.Pokemon.IsMythical.Set(fullPokemonSpecies.IsMythical),
						db.Pokemon.Color.Set(db.Color(fullPokemonSpecies.Color.Name)),
						db.Pokemon.Shape.Set(db.Shape(fullPokemonSpecies.Shape.Name)),
						db.Pokemon.IsDefaultVariant.Set(pokemon.IsDefault),
						db.Pokemon.Genus.Set(englishGenus.Genus),
						db.Pokemon.Height.Set(pokemon.Height),
						db.Pokemon.Weight.Set(pokemon.Weight),
						db.Pokemon.Habitat.SetIfPresent(habitat),
						db.Pokemon.Description.SetIfPresent(description),
						db.Pokemon.Sprite.SetIfPresent(&pokemon.Sprites.FrontDefault),
						db.Pokemon.EggGroups.Link(eggGroups...),
						db.Pokemon.UpdatedAt.Set(time.Now())).
					Exec(ctx)

				if dbErr != nil {
					log.Logger.WithField("pokemon", pokemon.Name).Fatal(dbErr)
				}

				for i := range pokemon.Types {
					t, _ := prisma.Type.FindUnique(db.Type.Slug.Equals(pokemon.Types[i].Type.Name)).Exec(ctx)
					_, err := prisma.PokemonType.UpsertOne(
						db.PokemonType.PokemonIDTypeID(db.PokemonType.PokemonID.Equals(createdPokemon.ID), db.PokemonType.TypeID.Equals(t.ID)),
					).Create(
						db.PokemonType.Pokemon.Link(db.Pokemon.Slug.Equals(pokemon.Name)),
						db.PokemonType.Type.Link(db.Type.Slug.Equals(pokemon.Types[i].Type.Name)),
						db.PokemonType.Slot.Set(pokemon.Types[i].Slot),
					).Update(
						db.PokemonType.Slot.Set(pokemon.Types[i].Slot),
					).Exec(ctx)

					if err != nil {
						log.Logger.WithField("pokemon", pokemon.Name).WithField("type", pokemon.Types[i].Type.Name).Fatal(dbErr)
					}
				}

				for i := range pokemon.Abilities {
					abilityName := pokemon.Abilities[i].Ability.Name

					if abilityName == "as-one" && pokemon.Name == "calyrex-shadow-rider" {
						abilityName = "as-one-shadow-rider"
					}

					if abilityName == "as-one" && pokemon.Name == "calyrex-ice-rider" {
						abilityName = "as-one-ice-rider"
					}

					ability, _ := prisma.Ability.FindUnique(db.Ability.Slug.Equals(abilityName)).Exec(ctx)
					_, err := prisma.PokemonAbility.UpsertOne(
						db.PokemonAbility.PokemonIDAbilityID(db.PokemonAbility.PokemonID.Equals(createdPokemon.ID), db.PokemonAbility.AbilityID.Equals(ability.ID)),
					).Create(
						db.PokemonAbility.Pokemon.Link(db.Pokemon.Slug.Equals(pokemon.Name)),
						db.PokemonAbility.Ability.Link(db.Ability.Slug.Equals(abilityName)),
						db.PokemonAbility.Slot.Set(pokemon.Abilities[i].Slot),
						db.PokemonAbility.IsHidden.Set(pokemon.Abilities[i].IsHidden),
					).Update(
						db.PokemonAbility.Slot.Set(pokemon.Abilities[i].Slot),
					).Exec(ctx)

					if err != nil {
						log.Logger.WithField("pokemon", pokemon.Name).WithField("ability", abilityName).Fatal(err)
					}
				}

				for i := range pokemon.Moves {
					for _, versionGroup := range pokemon.Moves[i].VersionGroupDetails {
						if versionGroup.VersionGroup.Name == "sword-shield" && !strings.Contains(pokemon.Moves[i].Move.Name, "max-") {
							m, _ := prisma.Move.FindUnique(db.Move.Slug.Equals(pokemon.Moves[i].Move.Name)).Exec(ctx)
							_, err := prisma.PokemonMove.UpsertOne(
								db.PokemonMove.PokemonIDMoveIDLearnMethod(
									db.PokemonMove.PokemonID.Equals(createdPokemon.ID),
									db.PokemonMove.MoveID.Equals(m.ID),
									db.PokemonMove.LearnMethod.Equals(db.MoveLearnMethod(versionGroup.MoveLearnMethod.Name)),
								),
							).Create(
								db.PokemonMove.Pokemon.Link(db.Pokemon.Slug.Equals(pokemon.Name)),
								db.PokemonMove.Move.Link(db.Move.Slug.Equals(pokemon.Moves[i].Move.Name)),
								db.PokemonMove.LearnMethod.Set(db.MoveLearnMethod(versionGroup.MoveLearnMethod.Name)),
								db.PokemonMove.LevelLearnedAt.Set(versionGroup.LevelLearnedAt),
							).Update(
								db.PokemonMove.LevelLearnedAt.Set(versionGroup.LevelLearnedAt),
							).Exec(ctx)

							if err != nil {
								log.Logger.WithField("pokemon", pokemon.Name).WithField("move", pokemon.Moves[i].Move.Name).Fatal(err)
							}
						}
					}
				}

				evolutionUrlParts := strings.Split(fullPokemonSpecies.EvolutionChain.Url, "/")
				evolutionChainId := evolutionUrlParts[len(evolutionUrlParts)-2]
				pokemonEvolutionChainList = append(pokemonEvolutionChainList, evolutionChainId)
			}
		}(i)
	}

	wg.Wait()

	pokemonEvolutionChainList = unique(pokemonEvolutionChainList)
	pokemonEvolutionChainListLength := len(pokemonEvolutionChainList)
	var wg2 sync.WaitGroup
	wg2.Add(pokemonEvolutionChainListLength)
	sem2 := make(chan bool, 20)

	for i := 0; i < pokemonEvolutionChainListLength; i++ {
		sem2 <- true
		go func(i int) {
			defer func() { <-sem2 }()
			defer wg2.Done()

			evolutionChain := client.GetEvolutionChain(pokemonEvolutionChainList[i])
			for _, chain := range evolutionChain.Chain.EvolvesTo {
				errors := insertEvolution(evolutionChain.Chain, chain, prisma, &ctx)

				for _, err := range errors {
					log.Logger.WithField("pokemon", chain.Species.Name).Fatal(err)
				}
			}
		}(i)
	}

	wg2.Wait()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Completed in %s\n", elapsed))
}

func insertEvolution(evolution pokeapi.Evolution, chain pokeapi.Evolution, prisma *db.PrismaClient, ctx *context.Context) []error {
	errors := make([]error, 0)
	fromPokemon, err := prisma.Pokemon.FindUnique(db.Pokemon.Slug.Equals(evolution.Species.Name)).Exec(*ctx)
	if err != nil {
		errors = append(errors, err)
	}

	for _, details := range chain.EvolutionDetails {
		if details.Location == nil && (evolution.Species.Name != "feebas" || (evolution.Species.Name == "feebas" && details.Trigger.Name == "trade")) {
			timeOfDay := db.TimeOfDay(*details.TimeOfDay)
			trigger := db.EvolutionTrigger(details.Trigger.Name)

			var item *string
			if details.Item != nil {
				item = &details.Item.Name
			}
			var heldItem *string
			if details.HeldItem != nil {
				heldItem = &details.HeldItem.Name
			}
			var knownMove *string
			if details.KnownMove != nil {
				knownMove = &details.KnownMove.Name
			}
			var knownMoveType *string
			if details.KnownMoveType != nil {
				knownMoveType = &details.KnownMoveType.Name
			}
			var partySpecies *string
			if details.PartySpecies != nil {
				partySpecies = &details.PartySpecies.Name
			}
			var partyType *string
			if details.PartyType != nil {
				partyType = &details.PartyType.Name
			}
			var tradeSpecies *string
			if details.TradeSpecies != nil {
				tradeSpecies = &details.TradeSpecies.Name
			}

			var gender db.Gender
			if details.Gender != nil {
				switch *details.Gender {
				case 1:
					gender = db.GenderMALE
				case 2:
					gender = db.GenderFEMALE
				default:
					gender = db.GenderANY
				}
			} else {
				gender = db.GenderANY
			}

			toPokemonSlug := chain.Species.Name
			if toPokemonSlug == "meowstic" {
				toPokemonSlug = "meowstic-male"

				if details.Gender != nil && gender == db.GenderFEMALE {
					toPokemonSlug = "meowstic-female"
				}
			}
			if toPokemonSlug == "toxtricity" {
				toPokemonSlug = "toxtricity-amped"
			}
			if toPokemonSlug == "urshifu" {
				toPokemonSlug = "urshifu-single-strike"
			}
			if toPokemonSlug == "lycanroc" {
				toPokemonSlug = "lycanroc-midday"

				if details.TimeOfDay != nil && db.TimeOfDay(*details.TimeOfDay) == db.TimeOfDayNIGHT {
					toPokemonSlug = "lycanroc-midnight"
				}
			}

			multipleSlugs := []string{toPokemonSlug}
			if toPokemonSlug == "toxtricity-amped" {
				multipleSlugs = append(multipleSlugs, "toxtricity-low-key")
			}
			if toPokemonSlug == "urshifu-single-strike" {
				multipleSlugs = append(multipleSlugs, "urshifu-rapid-strike")
			}

			for _, slug := range multipleSlugs {
				toPokemon, err := prisma.Pokemon.FindUnique(db.Pokemon.Slug.Equals(slug)).Exec(*ctx)
				if err != nil {
					errors = append(errors, err)
				}
				_, dbErr := prisma.PokemonEvolution.
					UpsertOne(db.PokemonEvolution.FromPokemonIDToPokemonIDTimeOfDayGenderTrigger(
						db.PokemonEvolution.FromPokemonID.Equals(fromPokemon.ID),
						db.PokemonEvolution.ToPokemonID.Equals(toPokemon.ID),
						db.PokemonEvolution.TimeOfDay.Equals(db.TimeOfDay(timeOfDay)),
						db.PokemonEvolution.Gender.Equals(gender),
						db.PokemonEvolution.Trigger.Equals(trigger),
					)).
					Create(
						db.PokemonEvolution.FromPokemon.Link(db.Pokemon.ID.Equals(fromPokemon.ID)),
						db.PokemonEvolution.ToPokemon.Link(db.Pokemon.ID.Equals(toPokemon.ID)),
						db.PokemonEvolution.Trigger.Set(trigger),
						db.PokemonEvolution.Gender.Set(gender),
						db.PokemonEvolution.NeedsOverworldRain.Set(details.NeedsOverworldRain),
						db.PokemonEvolution.TimeOfDay.Set(timeOfDay),
						db.PokemonEvolution.TurnUpsideDown.Set(details.TurnUpsideDown),
						db.PokemonEvolution.Spin.Set(false),
						db.PokemonEvolution.MinLevel.SetIfPresent(&details.MinLevel),
						db.PokemonEvolution.MinHappiness.SetIfPresent(&details.MinHappiness),
						db.PokemonEvolution.MinBeauty.SetIfPresent(&details.MinBeauty),
						db.PokemonEvolution.MinAffection.SetIfPresent(&details.MinAffection),
						db.PokemonEvolution.RelativePhysicalStats.SetIfPresent(&details.RelativePhysicalStats),
						db.PokemonEvolution.Item.Link(db.Item.Slug.EqualsIfPresent(item)),
						db.PokemonEvolution.HeldItem.Link(db.Item.Slug.EqualsIfPresent(heldItem)),
						db.PokemonEvolution.KnownMove.Link(db.Move.Slug.EqualsIfPresent(knownMove)),
						db.PokemonEvolution.KnownMoveType.Link(db.Type.Slug.EqualsIfPresent(knownMoveType)),
						db.PokemonEvolution.PartyPokemon.Link(db.Pokemon.Slug.EqualsIfPresent(partySpecies)),
						db.PokemonEvolution.PartyType.Link(db.Type.Slug.EqualsIfPresent(partyType)),
						db.PokemonEvolution.TradeWithPokemon.Link(db.Pokemon.Slug.EqualsIfPresent(tradeSpecies)),
					).
					Update(
						db.PokemonEvolution.NeedsOverworldRain.Set(details.NeedsOverworldRain),
						db.PokemonEvolution.TurnUpsideDown.Set(details.TurnUpsideDown),
						db.PokemonEvolution.Spin.Set(false),
						db.PokemonEvolution.MinLevel.SetIfPresent(&details.MinLevel),
						db.PokemonEvolution.MinHappiness.SetIfPresent(&details.MinHappiness),
						db.PokemonEvolution.MinBeauty.SetIfPresent(&details.MinBeauty),
						db.PokemonEvolution.MinAffection.SetIfPresent(&details.MinAffection),
						db.PokemonEvolution.RelativePhysicalStats.SetIfPresent(&details.RelativePhysicalStats),
						db.PokemonEvolution.Item.Link(db.Item.Slug.EqualsIfPresent(item)),
						db.PokemonEvolution.HeldItem.Link(db.Item.Slug.EqualsIfPresent(heldItem)),
						db.PokemonEvolution.KnownMove.Link(db.Move.Slug.EqualsIfPresent(knownMove)),
						db.PokemonEvolution.KnownMoveType.Link(db.Type.Slug.EqualsIfPresent(knownMoveType)),
						db.PokemonEvolution.PartyPokemon.Link(db.Pokemon.Slug.EqualsIfPresent(partySpecies)),
						db.PokemonEvolution.PartyType.Link(db.Type.Slug.EqualsIfPresent(partyType)),
						db.PokemonEvolution.TradeWithPokemon.Link(db.Pokemon.Slug.EqualsIfPresent(tradeSpecies)),
					).
					Exec(*ctx)

				if dbErr != nil {
					errors = append(errors, dbErr)
				}
			}

			for _, innerChain := range chain.EvolvesTo {
				errs := insertEvolution(chain, innerChain, prisma, ctx)
				if len(errors) != 0 {
					errors = append(errors, errs...)
				}
			}
		}
	}

	return errors
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
