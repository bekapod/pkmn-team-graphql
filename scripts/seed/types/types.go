package main

import (
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/scripts"
	"fmt"
	"strings"
	"sync"

	"github.com/alexflint/go-arg"
)

type RawDamageRelations struct {
	NoDamageTo       []*scripts.ResourcePointer `json:"no_damage_to"`
	HalfDamageTo     []*scripts.ResourcePointer `json:"half_damage_to"`
	DoubleDamageTo   []*scripts.ResourcePointer `json:"double_damage_to"`
	NoDamageFrom     []*scripts.ResourcePointer `json:"no_damage_from"`
	HalfDamageFrom   []*scripts.ResourcePointer `json:"half_damage_from"`
	DoubleDamageFrom []*scripts.ResourcePointer `json:"double_damage_from"`
}

type RawTypeList struct {
	Count   int                        `json:"count"`
	Results []*scripts.ResourcePointer `json:"results"`
}

type RawPokemon struct {
	Slot    int                     `json:"slot"`
	Pokemon scripts.ResourcePointer `json:"pokemon"`
}

type RawFullType struct {
	Id              int                          `json:"id"`
	Name            string                       `json:"name"`
	DamageRelations RawDamageRelations           `json:"damage_relations"`
	Names           []*scripts.RawTranslatedName `json:"names"`
	Pokemon         []*RawPokemon                `json:"pokemon"`
	Moves           []*scripts.ResourcePointer   `json:"moves"`
}

func getAllTypes() RawTypeList {
	typeList := RawTypeList{}
	scripts.GetResource("type", &typeList)
	return typeList
}

func getFullType(id string) RawFullType {
	fullType := RawFullType{}
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
			englishName, err := scripts.GetEnglishName(fullType.Names, fullType.Name)
			scripts.Check(err)
			values = append(values, fmt.Sprintf("('%s', '%s')", fullType.Name, englishName.Name))
		}(i)
	}

	wg.Wait()

	return fmt.Sprintf("INSERT INTO types (slug, name)\n\tVALUES %s;", strings.Join(values, ", "))
}

func main() {
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

	log.Logger.Info(fmt.Sprintf("Wrote %d bytes\n", o))
}
