package scripts

import (
	"bekapod/pkmn-team-graphql/log"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type SeedArgs struct {
	OutputFile  string `arg:"--output-file,env:OUTPUT_FILE"`
	PokeApiHost string `arg:"--pokeapi-host,env:POKEAPI_HOST" default:"https://pokeapi.co"`
}

type ResourcePointer struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type RawTranslatedName struct {
	Name     string          `json:"name"`
	Language ResourcePointer `json:"language"`
}

var Args *SeedArgs

var PokeApiPrefix string = "api/v2"

var Client = &http.Client{Timeout: 10 * time.Second}

func Check(e error) {
	if e != nil {
		log.Logger.Fatal(e)
	}
}

func OpenFile(path string) *os.File {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path := strings.Split(path, "/")
		dir := strings.Join(path[:len(path)-1], "/")
		log.Logger.Debug(fmt.Sprintf("Creating directory: %s", dir))
		os.MkdirAll(dir, 0700) // Create your file
	}

	f, err := os.Create(path)
	Check(err)

	return f
}

func GetResource(path string, bucket interface{}) {
	resourceUrl := fmt.Sprintf("%s/%s/%s", Args.PokeApiHost, PokeApiPrefix, path)
	log.Logger.WithField("resource", resourceUrl).Info("Getting resource")
	response, err := Client.Get(resourceUrl)
	Check(err)
	log.Logger.WithField("resource", resourceUrl).Info("Got resource")
	defer response.Body.Close()
	decodeErr := json.NewDecoder(response.Body).Decode(bucket)
	Check(decodeErr)
}

func GetEnglishName(names []*RawTranslatedName, resourceName string) (*RawTranslatedName, error) {
	for i := range names {
		if names[i].Language.Name == "en" {
			return names[i], nil
		}
	}

	return &RawTranslatedName{}, fmt.Errorf("no english name found for %s", resourceName)
}
