package main

import (
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/scripts"
	"bekapod/pkmn-team-graphql/scripts/pokeapi"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
)

func getAllMoves() pokeapi.ResourcePointerList {
	moveList := pokeapi.ResourcePointerList{}
	scripts.GetResource("move?limit=1000", &moveList)
	return moveList
}

func getMove(name string) pokeapi.RawMove {
	fullMove := pokeapi.RawMove{}
	scripts.GetResource(fmt.Sprintf("move/%s", name), &fullMove)
	return fullMove
}

func getAllMoveTargets() pokeapi.ResourcePointerList {
	moveTargetList := pokeapi.ResourcePointerList{}
	scripts.GetResource("move-target", &moveTargetList)
	return moveTargetList
}

func getMoveTarget(name string) pokeapi.RawTarget {
	target := pokeapi.RawTarget{}
	scripts.GetResource(fmt.Sprintf("move-target/%s", name), &target)
	return target
}

func getListOfMoveTargets() map[string]string {
	moveTargetList := getAllMoveTargets()
	results := moveTargetList.Results
	resultsLength := len(results)
	targets := make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			target := getMoveTarget(results[i].Name)
			englishName, err := pokeapi.GetEnglishName(target.Names, target.Name)
			scripts.Check(err)
			targets[target.Name] = englishName.Name
		}(i)
	}

	wg.Wait()
	return targets
}

func generateMovesSeed() string {
	moveList := getAllMoves()
	targets := getListOfMoveTargets()
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
			fullMove := getMove(results[i].Name)
			if fullMove.DamageClass.Name != "" {
				target := targets[fullMove.Target.Name]

				englishName, err := pokeapi.GetEnglishName(fullMove.Names, fullMove.Name)
				scripts.Check(err)
				englishEffectEntry, _ := pokeapi.GetEnglishEffectEntry(fullMove.EffectEntries, fullMove.Name)

				values = append(values, fmt.Sprintf(
					"('%s', '%s', %d, %d, %d, '%s', '%s', %d, '%s', %s)",
					scripts.EscapeSingleQuote(fullMove.Name),
					scripts.EscapeSingleQuote(englishName.Name),
					fullMove.Accuracy,
					fullMove.PP,
					fullMove.Power,
					scripts.EscapeSingleQuote(fullMove.DamageClass.Name),
					scripts.EscapeSingleQuote(englishEffectEntry.ShortEffect),
					fullMove.EffectChance,
					scripts.EscapeSingleQuote(target),
					fmt.Sprintf("(SELECT id from types WHERE slug='%s')", fullMove.Type.Name),
				))
			}
		}(i)
	}

	wg.Wait()

	return fmt.Sprintf("INSERT INTO moves (slug, name, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id)\n\tVALUES %s;", strings.Join(values, ", "))
}

func main() {
	start := time.Now()
	scripts.Args = &scripts.SeedArgs{
		OutputFile: "seeds/moves.sql",
	}
	arg.MustParse(scripts.Args)

	f := scripts.OpenFile(scripts.Args.OutputFile)
	defer f.Close()
	sql := generateMovesSeed()

	o, err := f.WriteString(sql)
	scripts.Check(err)
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
