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

func getAllStats() pokeapi.ResourcePointerList {
	statList := pokeapi.ResourcePointerList{}
	scripts.GetResource("stat?limit=1000", &statList)
	return statList
}

func getStat(name string) pokeapi.RawStat {
	fullStat := pokeapi.RawStat{}
	scripts.GetResource(fmt.Sprintf("stat/%s", name), &fullStat)
	return fullStat
}

func generateStatsSeed() string {
	statList := getAllStats()
	results := statList.Results
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
			fullStat := getStat(results[i].Name)
			englishName, err := pokeapi.GetEnglishName(fullStat.Names, fullStat.Name)
			scripts.Check(err)

			values = append(values, fmt.Sprintf(
				"('%s', '%s')",
				scripts.EscapeSingleQuote(fullStat.Name),
				scripts.EscapeSingleQuote(englishName.Name),
			))
		}(i)
	}

	wg.Wait()

	return fmt.Sprintf("INSERT INTO stats (slug, name)\n\tVALUES %s;", strings.Join(values, ", "))
}

func main() {
	start := time.Now()
	scripts.Args = &scripts.SeedArgs{
		OutputFile: "seeds/stats.sql",
	}
	arg.MustParse(scripts.Args)

	f := scripts.OpenFile(scripts.Args.OutputFile)
	defer f.Close()
	sql := generateStatsSeed()

	o, err := f.WriteString(sql)
	scripts.Check(err)
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
