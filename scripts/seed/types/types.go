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

func getAllTypes() pokeapi.ResourcePointerList {
	typeList := pokeapi.ResourcePointerList{}
	scripts.GetResource("type", &typeList)
	return typeList
}

func getFullType(id string) pokeapi.RawType {
	fullType := pokeapi.RawType{}
	scripts.GetResource(fmt.Sprintf("type/%s", id), &fullType)
	return fullType
}

func generateTypesSeed() string {
	typeList := getAllTypes()
	results := typeList.Results
	resultsLength := len(results)
	values := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(resultsLength)

	for i := 0; i < resultsLength; i++ {
		go func(i int) {
			defer wg.Done()
			urlParts := strings.Split(results[i].Url, "/")
			id := urlParts[len(urlParts)-2]
			fullType := getFullType(id)
			englishName, err := pokeapi.GetEnglishName(fullType.Names, fullType.Name)
			scripts.Check(err)
			values = append(values, fmt.Sprintf(
				"('%s', '%s')",
				scripts.EscapeSingleQuote(fullType.Name),
				scripts.EscapeSingleQuote(englishName.Name),
			))
		}(i)
	}

	wg.Wait()

	return fmt.Sprintf("INSERT INTO types (slug, name)\n\tVALUES %s;", strings.Join(values, ", "))
}

func main() {
	start := time.Now()
	scripts.Args = &scripts.SeedArgs{
		OutputFile: "seeds/types.sql",
	}
	arg.MustParse(scripts.Args)

	f := scripts.OpenFile(scripts.Args.OutputFile)
	defer f.Close()
	sql := generateTypesSeed()

	o, err := f.WriteString(sql)
	scripts.Check(err)
	f.Sync()

	elapsed := time.Since(start)
	log.Logger.Info(fmt.Sprintf("Wrote %d bytes in %s\n", o, elapsed))
}
