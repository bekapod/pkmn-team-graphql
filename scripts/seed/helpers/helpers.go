package helpers

import (
	"bekapod/pkmn-team-graphql/log"
	"bekapod/pkmn-team-graphql/pokeapi"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	pokeapi.PokeApiConfig
	OutputFile string `arg:"--output-file,env:OUTPUT_FILE"`
}

func OpenFile(path string) *os.File {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path := strings.Split(path, "/")
		dir := strings.Join(path[:len(path)-1], "/")
		log.Logger.Debug(fmt.Sprintf("Creating directory: %s", dir))
		os.MkdirAll(dir, 0700) // Create your file
	}

	f, err := os.Create(path)
	if err != nil {
		log.Logger.Fatal(err)
	}

	return f
}

func EscapeSingleQuote(str string) string {
	return strings.Replace(str, "'", "''", -1)
}

func Chunk(numberOfChunks int, items []string) [][]string {
	var chunked [][]string
	chunkSize := (len(items) + numberOfChunks - 1) / numberOfChunks

	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize

		if end > len(items) {
			end = len(items)
		}

		chunked = append(chunked, items[i:end])
	}

	return chunked
}
