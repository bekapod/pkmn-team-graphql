package pokeapi

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/log"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PokeApiConfig struct {
	Host   string `arg:"--pokeapi-host,env:POKEAPI_HOST" default:"https://pokeapi.co"`
	Prefix string `default:"api/v2"`
}

type PokeApiClient struct {
	httpClient *http.Client
	config     *PokeApiConfig
}

func New(config PokeApiConfig) *PokeApiClient {
	return &PokeApiClient{
		httpClient: &http.Client{Timeout: 60 * time.Second},
		config:     &config,
	}
}

func (client PokeApiClient) GetResourceList(resource string, limit int) ResourcePointerList {
	result := ResourcePointerList{}
	client.getResource(fmt.Sprintf("%s/?limit=%d", resource, limit), &result)
	return result
}

func (client PokeApiClient) GetAbility(idOrName string) Ability {
	result := Ability{}
	client.getResource(fmt.Sprintf("ability/%s/", idOrName), &result)
	return result
}

func (client PokeApiClient) GetEggGroup(idOrName string) EggGroup {
	result := EggGroup{}
	client.getResource(fmt.Sprintf("egg-group/%s/", idOrName), &result)
	return result
}

func (client PokeApiClient) GetEvolutionChain(id string) EvolutionChain {
	result := EvolutionChain{}
	client.getResource(fmt.Sprintf("evolution-chain/%s/", id), &result)
	normalizeEvolutionChain(&result.Chain)
	return result
}

func normalizeEvolutionChain(chain *Evolution) {
	for i := range chain.EvolutionDetails {
		switch chain.EvolutionDetails[i].Trigger.Name {
		case "level-up":
			chain.EvolutionDetails[i].Trigger.Name = string(db.EvolutionTriggerLEVELUP)
		case "other":
			chain.EvolutionDetails[i].Trigger.Name = string(db.EvolutionTriggerOTHER)
		case "shed":
			chain.EvolutionDetails[i].Trigger.Name = string(db.EvolutionTriggerSHED)
		case "trade":
			chain.EvolutionDetails[i].Trigger.Name = string(db.EvolutionTriggerTRADE)
		case "use-item":
			chain.EvolutionDetails[i].Trigger.Name = string(db.EvolutionTriggerUSEITEM)
		}

		day := string(db.TimeOfDayDAY)
		night := string(db.TimeOfDayNIGHT)
		anyTime := string(db.TimeOfDayANY)

		if chain.EvolutionDetails[i].TimeOfDay != nil {
			switch *chain.EvolutionDetails[i].TimeOfDay {
			case "day":
				chain.EvolutionDetails[i].TimeOfDay = &day
			case "night":
				chain.EvolutionDetails[i].TimeOfDay = &night
			default:
				chain.EvolutionDetails[i].TimeOfDay = &anyTime
			}
		} else {
			chain.EvolutionDetails[i].TimeOfDay = &anyTime
		}
	}

	for i := range chain.EvolvesTo {
		normalizeEvolutionChain(&chain.EvolvesTo[i])
	}
}

func (client PokeApiClient) GetGeneration(idOrName string) Generation {
	result := Generation{}
	client.getResource(fmt.Sprintf("generation/%s/", idOrName), &result)
	return result
}

func (client PokeApiClient) GetItem(nameOrId string) Item {
	result := Item{}
	client.getResource(fmt.Sprintf("item/%s/", nameOrId), &result)

	switch result.Category.Name {
	case "all-machines":
		result.Category.Name = string(db.ItemCategoryALLMACHINES)
	case "all-mail":
		result.Category.Name = string(db.ItemCategoryALLMAIL)
	case "apricorn-balls":
		result.Category.Name = string(db.ItemCategoryAPRICORNBALLS)
	case "apricorn-box":
		result.Category.Name = string(db.ItemCategoryAPRICORNBOX)
	case "bad-held-items":
		result.Category.Name = string(db.ItemCategoryBADHELDITEMS)
	case "baking-only":
		result.Category.Name = string(db.ItemCategoryBAKINGONLY)
	case "choice":
		result.Category.Name = string(db.ItemCategoryCHOICE)
	case "collectibles":
		result.Category.Name = string(db.ItemCategoryCOLLECTIBLES)
	case "data-cards":
		result.Category.Name = string(db.ItemCategoryDATACARDS)
	case "dex-completion":
		result.Category.Name = string(db.ItemCategoryDEXCOMPLETION)
	case "effort-drop":
		result.Category.Name = string(db.ItemCategoryEFFORTDROP)
	case "effort-training":
		result.Category.Name = string(db.ItemCategoryEFFORTTRAINING)
	case "event-items":
		result.Category.Name = string(db.ItemCategoryEVENTITEMS)
	case "evolution":
		result.Category.Name = string(db.ItemCategoryEVOLUTION)
	case "flutes":
		result.Category.Name = string(db.ItemCategoryFLUTES)
	case "gameplay":
		result.Category.Name = string(db.ItemCategoryGAMEPLAY)
	case "healing":
		result.Category.Name = string(db.ItemCategoryHEALING)
	case "held-items":
		result.Category.Name = string(db.ItemCategoryHELDITEMS)
	case "in-a-pinch":
		result.Category.Name = string(db.ItemCategoryINAPINCH)
	case "jewels":
		result.Category.Name = string(db.ItemCategoryJEWELS)
	case "loot":
		result.Category.Name = string(db.ItemCategoryLOOT)
	case "medicine":
		result.Category.Name = string(db.ItemCategoryMEDICINE)
	case "mega-stones":
		result.Category.Name = string(db.ItemCategoryMEGASTONES)
	case "memories":
		result.Category.Name = string(db.ItemCategoryMEMORIES)
	case "miracle-shooter":
		result.Category.Name = string(db.ItemCategoryMIRACLESHOOTER)
	case "mulch":
		result.Category.Name = string(db.ItemCategoryMULCH)
	case "other":
		result.Category.Name = string(db.ItemCategoryOTHER)
	case "picky-healing":
		result.Category.Name = string(db.ItemCategoryPICKYHEALING)
	case "plates":
		result.Category.Name = string(db.ItemCategoryPLATES)
	case "plot-advancement":
		result.Category.Name = string(db.ItemCategoryPLOTADVANCEMENT)
	case "pp-recovery":
		result.Category.Name = string(db.ItemCategoryPPRECOVERY)
	case "revival":
		result.Category.Name = string(db.ItemCategoryREVIVAL)
	case "scarves":
		result.Category.Name = string(db.ItemCategorySCARVES)
	case "special-balls":
		result.Category.Name = string(db.ItemCategorySPECIALBALLS)
	case "species-specific":
		result.Category.Name = string(db.ItemCategorySPECIESSPECIFIC)
	case "spelunking":
		result.Category.Name = string(db.ItemCategorySPELUNKING)
	case "standard-balls":
		result.Category.Name = string(db.ItemCategorySTANDARDBALLS)
	case "stat-boosts":
		result.Category.Name = string(db.ItemCategorySTATBOOSTS)
	case "status-cures":
		result.Category.Name = string(db.ItemCategorySTATUSCURES)
	case "training":
		result.Category.Name = string(db.ItemCategoryTRAINING)
	case "type-enhancement":
		result.Category.Name = string(db.ItemCategoryTYPEENHANCEMENT)
	case "type-protection":
		result.Category.Name = string(db.ItemCategoryTYPEPROTECTION)
	case "unused":
		result.Category.Name = string(db.ItemCategoryUNUSED)
	case "vitamins":
		result.Category.Name = string(db.ItemCategoryVITAMINS)
	case "z-crystals":
		result.Category.Name = string(db.ItemCategoryZCRYSTALS)
	}

	for i, attribute := range result.Attributes {
		switch attribute.Name {
		case "consumable":
			result.Attributes[i].Name = string(db.ItemAttributeCONSUMABLE)
		case "countable":
			result.Attributes[i].Name = string(db.ItemAttributeCOUNTABLE)
		case "holdable":
			result.Attributes[i].Name = string(db.ItemAttributeHOLDABLE)
		case "holdable-active":
			result.Attributes[i].Name = string(db.ItemAttributeHOLDABLEACTIVE)
		case "holdable-passive":
			result.Attributes[i].Name = string(db.ItemAttributeHOLDABLEPASSIVE)
		case "underground":
			result.Attributes[i].Name = string(db.ItemAttributeUNDERGROUND)
		case "usable-in-battle":
			result.Attributes[i].Name = string(db.ItemAttributeUSABLEINBATTLE)
		case "usable-overworld":
			result.Attributes[i].Name = string(db.ItemAttributeUSABLEOVERWORLD)
		}
	}

	return result
}

func (client PokeApiClient) GetLocation(nameOrId string) Location {
	result := Location{}
	client.getResource(fmt.Sprintf("location/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetMove(idOrName string) Move {
	result := Move{}
	client.getResource(fmt.Sprintf("move/%s/", idOrName), &result)

	switch result.DamageClass.Name {
	case "physical":
		result.DamageClass.Name = string(db.DamageClassPHYSICAL)
	case "special":
		result.DamageClass.Name = string(db.DamageClassSPECIAL)
	case "status":
		result.DamageClass.Name = string(db.DamageClassSTATUS)
	}

	switch result.Target.Name {
	case "specific-move":
		result.Target.Name = string(db.MoveTargetSPECIFICMOVE)
	case "selected-pokemon-me-first":
		result.Target.Name = string(db.MoveTargetSELECTEDPOKEMONMEFIRST)
	case "ally":
		result.Target.Name = string(db.MoveTargetALLY)
	case "users-field":
		result.Target.Name = string(db.MoveTargetUSERSFIELD)
	case "user-or-ally":
		result.Target.Name = string(db.MoveTargetUSERORALLY)
	case "opponents-field":
		result.Target.Name = string(db.MoveTargetOPPONENTSFIELD)
	case "user":
		result.Target.Name = string(db.MoveTargetUSER)
	case "random-opponent":
		result.Target.Name = string(db.MoveTargetRANDOMOPPONENT)
	case "all-other-pokemon":
		result.Target.Name = string(db.MoveTargetALLOTHERPOKEMON)
	case "selected-pokemon":
		result.Target.Name = string(db.MoveTargetSELECTEDPOKEMON)
	case "all-opponents":
		result.Target.Name = string(db.MoveTargetALLOPPONENTS)
	case "entire-field":
		result.Target.Name = string(db.MoveTargetENTIREFIELD)
	case "user-and-allies":
		result.Target.Name = string(db.MoveTargetUSERANDALLIES)
	case "all-pokemon":
		result.Target.Name = string(db.MoveTargetALLPOKEMON)
	case "all-allies":
		result.Target.Name = string(db.MoveTargetALLALLIES)
	}

	return result
}

func (client PokeApiClient) GetMoveTarget(nameOrId string) Target {
	result := Target{}
	client.getResource(fmt.Sprintf("move-target/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetPokedex(nameOrId string) Pokedex {
	result := Pokedex{}
	client.getResource(fmt.Sprintf("pokedex/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetPokemon(nameOrId string) Pokemon {
	result := Pokemon{}
	client.getResource(fmt.Sprintf("pokemon/%s/", nameOrId), &result)
	moves := make([]PokemonMove, 0)

	for _, move := range result.Moves {
		for i, versionGroup := range move.VersionGroupDetails {
			if versionGroup.VersionGroup.Name == "sword-shield" && !strings.Contains(move.Move.Name, "max-") {
				switch move.VersionGroupDetails[i].MoveLearnMethod.Name {
				case "level-up":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodLEVELUP)
				case "egg":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodEGG)
				case "tutor":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodTUTOR)
				case "machine":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodMACHINE)
				case "stadium-surfing-pikachu":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodSTADIUMSURFINGPIKACHU)
				case "light-ball-egg":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodLIGHTBALLEGG)
				case "colosseum-purification":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodCOLOSSEUMPURIFICATION)
				case "xd-shadow":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodXDSHADOW)
				case "xd-purification":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodXDPURIFICATION)
				case "form-change":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodFORMCHANGE)
				case "record":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodRECORD)
				case "transfer":
					move.VersionGroupDetails[i].MoveLearnMethod.Name = string(db.MoveLearnMethodTRANSFER)
				}
				moves = append(moves, move)
			}
		}
	}

	result.Moves = moves

	return result
}

func (client PokeApiClient) GetPokemonSpecies(nameOrId string) PokemonSpecies {
	result := PokemonSpecies{}
	client.getResource(fmt.Sprintf("pokemon-species/%s/", nameOrId), &result)

	switch result.Color.Name {
	case "black":
		result.Color.Name = string(db.ColorBLACK)
	case "blue":
		result.Color.Name = string(db.ColorBLUE)
	case "brown":
		result.Color.Name = string(db.ColorBROWN)
	case "gray":
		result.Color.Name = string(db.ColorGRAY)
	case "green":
		result.Color.Name = string(db.ColorGREEN)
	case "pink":
		result.Color.Name = string(db.ColorPINK)
	case "purple":
		result.Color.Name = string(db.ColorPURPLE)
	case "red":
		result.Color.Name = string(db.ColorRED)
	case "white":
		result.Color.Name = string(db.ColorWHITE)
	case "yellow":
		result.Color.Name = string(db.ColorYELLOW)
	}

	if result.Habitat != nil {
		switch result.Habitat.Name {
		case "cave":
			result.Habitat.Name = string(db.HabitatCAVE)
		case "forest":
			result.Habitat.Name = string(db.HabitatFOREST)
		case "grassland":
			result.Habitat.Name = string(db.HabitatGRASSLAND)
		case "mountain":
			result.Habitat.Name = string(db.HabitatMOUNTAIN)
		case "rare":
			result.Habitat.Name = string(db.HabitatRARE)
		case "rough-terrain":
			result.Habitat.Name = string(db.HabitatROUGHTERRAIN)
		case "sea":
			result.Habitat.Name = string(db.HabitatSEA)
		case "urban":
			result.Habitat.Name = string(db.HabitatURBAN)
		case "waters-edge":
			result.Habitat.Name = string(db.HabitatWATERSEDGE)
		}
	}

	switch result.Shape.Name {
	case "ball":
		result.Shape.Name = string(db.ShapeBALL)
	case "squiggle":
		result.Shape.Name = string(db.ShapeSQUIGGLE)
	case "fish":
		result.Shape.Name = string(db.ShapeFISH)
	case "arms":
		result.Shape.Name = string(db.ShapeARMS)
	case "blob":
		result.Shape.Name = string(db.ShapeBLOB)
	case "upright":
		result.Shape.Name = string(db.ShapeUPRIGHT)
	case "legs":
		result.Shape.Name = string(db.ShapeLEGS)
	case "quadruped":
		result.Shape.Name = string(db.ShapeQUADRUPED)
	case "wings":
		result.Shape.Name = string(db.ShapeWINGS)
	case "tentacles":
		result.Shape.Name = string(db.ShapeTENTACLES)
	case "heads":
		result.Shape.Name = string(db.ShapeHEADS)
	case "humanoid":
		result.Shape.Name = string(db.ShapeHUMANOID)
	case "bug-wings":
		result.Shape.Name = string(db.ShapeBUGWINGS)
	case "armor":
		result.Shape.Name = string(db.ShapeARMOR)
	}

	return result
}

func (client PokeApiClient) GetPokemonVarietiesForSpecies(species PokemonSpecies) []Pokemon {
	pokemon := make([]Pokemon, 0)

	for i := range species.Varieties {
		urlParts := strings.Split(species.Varieties[i].Pokemon.Url, "/")
		id := urlParts[len(urlParts)-2]
		variety := client.GetPokemon(id)
		excludeSlugPatterns := []string{
			"castform-",
			"-crowned",
			"deoxys-",
			"-eternamax",
			"gmax",
			"greninja-",
			"-mega",
			"necrozma-",
			"pikachu-",
			"pirouette",
			"-primal",
			"-resolute",
			"-therian",
			"totem",
			"zygarde-",
		}
		excludeSlugExact := []string{
			"aegislash-blade",
			"darmanitan-zen",
			"darmanitan-zen-galar",
			"eiscue-noice",
			"floette-eternal",
			"giratina-origin",
			"giratina-origin",
			"gourgeist-large",
			"gourgeist-small",
			"gourgeist-super",
			"hoopa-unbound",
			"magearna-original",
			"mimikyu-busted",
			"minior-green",
			"minior-indigo",
			"minior-orange",
			"minior-red",
			"minior-violet",
			"minior-yellow",
			"oricorio-pau",
			"oricorio-pom-pom",
			"oricorio-sensu",
			"pumpkaboo-large",
			"pumpkaboo-small",
			"pumpkaboo-super",
			"rockruff-own-tempo",
			"taillow",
			"wishiwashi-school",
			"wormadom-sandy",
			"wormadom-trash",
		}
		shouldInclude := true

		unnecessaryVariantNames := []string{
			"giratina-altered",
			"darmanitan-standard",
			"tornadus-incarnate",
			"thundurus-incarnate",
			"landorus-incarnate",
			"keldeo-ordinary",
			"aegislash-shield",
			"pumpkaboo-average",
			"gourgeist-average",
			"wishiwashi-solo",
			"mimikyu-disguised",
			"eiscue-ice",
			"zacian-hero",
			"zamazenta-hero",
		}

		for _, exact := range excludeSlugExact {
			if variety.Name == exact {
				shouldInclude = false
				break
			}
		}

		if shouldInclude {
			for _, pattern := range excludeSlugPatterns {
				if strings.Contains(variety.Name, pattern) {
					shouldInclude = false
					break
				}
			}
		}

		if shouldInclude {
			for _, pattern := range unnecessaryVariantNames {
				if variety.Name == pattern {
					variety.Name = variety.Species.Name
					break
				}
			}

			if variety.Name == "darmanitan-standard-galar" {
				variety.Name = "darmanitan-galar"
			}

			pokemon = append(pokemon, variety)
		}
	}

	return pokemon
}

func (client PokeApiClient) GetRegion(nameOrId string) Region {
	result := Region{}
	client.getResource(fmt.Sprintf("region/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetStat(nameOrId string) Stat {
	result := Stat{}
	client.getResource(fmt.Sprintf("stat/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetType(nameOrId string) Type {
	result := Type{}
	client.getResource(fmt.Sprintf("type/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) GetVersionGroup(nameOrId string) VersionGroup {
	result := VersionGroup{}
	client.getResource(fmt.Sprintf("version-group/%s/", nameOrId), &result)
	return result
}

func (client PokeApiClient) getResource(path string, bucket interface{}) {
	resourceUrl := fmt.Sprintf("%s/%s/%s", client.config.Host, client.config.Prefix, path)
	log.Logger.WithField("resource", resourceUrl).Info("Getting resource")
	response, err := client.httpClient.Get(resourceUrl)
	check(err, resourceUrl)
	log.Logger.WithField("resource", resourceUrl).Info("Got resource")
	defer response.Body.Close()
	decodeErr := json.NewDecoder(response.Body).Decode(bucket)
	check(decodeErr, resourceUrl)
}

func GetEnglishName(names []TranslatedName, resourceName string) (TranslatedName, error) {
	for i := range names {
		if names[i].Language.Name == "en" {
			return names[i], nil
		}
	}

	return TranslatedName{}, fmt.Errorf("no english name found for %s", resourceName)
}

func GetEnglishGenus(genera []Genus, resourceName string) (Genus, error) {
	for i := range genera {
		if genera[i].Language.Name == "en" {
			return genera[i], nil
		}
	}

	return Genus{
		Genus: "",
	}, fmt.Errorf("no english genus found for %s", resourceName)
}

func GetEnglishEffectEntry(effectEntries []EffectEntry, resourceName string) (*EffectEntry, error) {
	for i := range effectEntries {
		if effectEntries[i].Language.Name == "en" {
			return &effectEntries[i], nil
		}
	}

	return nil, fmt.Errorf("no english effect entry found for %s", resourceName)
}

func GetEnglishFlavourTextEntries(flavourTextEntries []FlavourTextEntry, resourceName string) (*[]FlavourTextEntry, error) {
	entries := make([]FlavourTextEntry, 0)
	for i := range flavourTextEntries {
		if flavourTextEntries[i].Language.Name == "en" && flavourTextEntries[i].Version.Name == "sword" {
			entries = append(entries, flavourTextEntries[i])
		}
	}

	if len(entries) > 0 {
		return &entries, nil
	}

	return nil, fmt.Errorf("no english flavour text entry found for %s", resourceName)
}

func GetPokemonStat(pkmn Pokemon, stat string) (int, error) {
	for i := range pkmn.Stats {
		if pkmn.Stats[i].Stat.Name == stat {
			return pkmn.Stats[i].BaseStat, nil
		}
	}

	return 0, fmt.Errorf("couldn't find %s stat for %s", stat, pkmn.Name)
}

func GetPokedexId(pkmn PokemonSpecies, pokedex string) (int, error) {
	for i := range pkmn.PokedexNumbers {
		if pkmn.PokedexNumbers[i].Pokedex.Name == pokedex {
			return pkmn.PokedexNumbers[i].EntryNumber, nil
		}
	}

	return 0, fmt.Errorf("couldn't find %s pokedex id for %s", pokedex, pkmn.Name)
}

func check(e error, resourceUrl string) {
	if e != nil {
		log.Logger.WithField("resource", resourceUrl).Fatal(e)
	}
}

type Ability struct {
	Id                 int                `json:"id"`
	Name               string             `json:"name"`
	IsMainSeries       bool               `json:"is_main_series"`
	Names              []TranslatedName   `json:"names"`
	EffectEntries      []EffectEntry      `json:"effect_entries"`
	FlavourTextEntries []FlavourTextEntry `json:"flavor_text_entries"`
	Pokemon            []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Pokemon  ResourcePointer
	} `json:"pokemon"`
}

type EffectEntry struct {
	Effect      string          `json:"effect"`
	ShortEffect string          `json:"short_effect"`
	Language    ResourcePointer `json:"language"`
}

type EggGroup struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Names          []TranslatedName  `json:"names"`
	PokemonSpecies []ResourcePointer `json:"pokemon_species"`
}

type Evolution struct {
	IsBaby           bool            `json:"is_baby"`
	Species          ResourcePointer `json:"species"`
	EvolvesTo        []Evolution     `json:"evolves_to"`
	EvolutionDetails []struct {
		Item                  *ResourcePointer `json:"item"`
		Trigger               ResourcePointer  `json:"trigger"`
		Gender                *int             `json:"gender"`
		HeldItem              *ResourcePointer `json:"held_item"`
		KnownMove             *ResourcePointer `json:"known_move"`
		KnownMoveType         *ResourcePointer `json:"known_move_type"`
		Location              *ResourcePointer `json:"location"`
		MinLevel              int              `json:"min_level"`
		MinHappiness          int              `json:"min_happiness"`
		MinBeauty             int              `json:"min_beauty"`
		MinAffection          int              `json:"min_affection"`
		NeedsOverworldRain    bool             `json:"needs_overworld_rain"`
		PartySpecies          *ResourcePointer `json:"party_species"`
		PartyType             *ResourcePointer `json:"party_type"`
		RelativePhysicalStats int              `json:"relative_physical_stats"`
		TimeOfDay             *string          `json:"time_of_day"`
		TradeSpecies          *ResourcePointer `json:"trade_species"`
		TurnUpsideDown        bool             `json:"turn_upside_down"`
	} `json:"evolution_details"`
}

type EvolutionChain struct {
	ID              int             `json:"id"`
	BabyTriggerItem ResourcePointer `json:"baby_trigger_item"`
	Chain           Evolution       `json:"chain"`
}

type FlavourTextEntry struct {
	FlavourText  string          `json:"flavor_text"`
	Language     ResourcePointer `json:"language"`
	VersionGroup ResourcePointer `json:"version_group"`
	Version      ResourcePointer `json:"version"`
}

type GameIndex struct {
	GameIndex  int             `json:"game_index"`
	Generation ResourcePointer `json:"generation"`
}

type Generation struct {
	Id             int               `json:"id"`
	Name           string            `json:"name"`
	Abilities      []ResourcePointer `json:"abilities"`
	MainRegion     ResourcePointer   `json:"main_region"`
	Moves          []ResourcePointer `json:"moves"`
	Names          []TranslatedName  `json:"names"`
	PokemonSpecies []ResourcePointer `json:"pokemon_species"`
	Types          []ResourcePointer `json:"types"`
	VersionGroups  []ResourcePointer `json:"version_groups"`
}

type Genus struct {
	Genus    string          `json:"genus"`
	Language ResourcePointer `json:"language"`
}

type Item struct {
	ID                 int                `json:"id"`
	Name               string             `json:"name"`
	Cost               *int               `json:"cost"`
	FlingPower         *int               `json:"fling_power"`
	FlingEffect        *ResourcePointer   `json:"fling_effect"`
	Attributes         []ResourcePointer  `json:"attributes"`
	Category           ResourcePointer    `json:"category"`
	EffectEntries      []EffectEntry      `json:"effect_entries"`
	FlavourTextEntries []FlavourTextEntry `json:"flavor_text_entries"`
	Names              []TranslatedName   `json:"names"`
	HeldByPokemon      []ResourcePointer  `json:"held_by_pokemon"`
	Sprites            struct {
		Default *string `json:"default"`
	} `json:"sprites"`
	GameIndices []GameIndex `json:"game_indices"`
}

type Location struct {
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	Region      ResourcePointer  `json:"region"`
	Names       []TranslatedName `json:"names"`
	GameIndices []GameIndex      `json:"game_indices"`
}

type Move struct {
	Id            int                `json:"id"`
	Name          string             `json:"name"`
	Accuracy      *int               `json:"accuracy"`
	EffectChance  *int               `json:"effect_chance"`
	PP            *int               `json:"pp"`
	Priority      int                `json:"priority"`
	Power         *int               `json:"power"`
	DamageClass   ResourcePointer    `json:"damage_class"`
	EffectEntries []EffectEntry      `json:"effect_entries"`
	Names         []TranslatedName   `json:"names"`
	Target        ResourcePointer    `json:"target"`
	Type          ResourcePointer    `json:"type"`
	Description   []FlavourTextEntry `json:"flavor_text_entries"`
}

type Pokedex struct {
	Id            int               `json:"id"`
	Name          string            `json:"name"`
	IsMainSeries  bool              `json:"is_main_series"`
	Names         []TranslatedName  `json:"names"`
	Region        ResourcePointer   `json:"region"`
	VersionGroups []ResourcePointer `json:"version_groups"`
	Descriptions  []struct {
		Description string          `json:"description"`
		Language    ResourcePointer `json:"language"`
	} `json:"descriptions"`
	PokemonEntries []struct {
		EntryNumber    int             `json:"entry_number"`
		PokemonSpecies ResourcePointer `json:"pokemon_species"`
	} `json:"pokemon_entries"`
}

type Pokemon struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	BaseExperience int               `json:"base_experience"`
	Height         int               `json:"height"`
	IsDefault      bool              `json:"is_default"`
	Order          int               `json:"order"`
	Weight         int               `json:"weight"`
	Forms          []ResourcePointer `json:"forms"`
	Species        ResourcePointer   `json:"species"`
	Abilities      []struct {
		IsHidden bool            `json:"is_hidden"`
		Slot     int             `json:"slot"`
		Ability  ResourcePointer `json:"ability"`
	} `json:"abilities"`
	Moves   []PokemonMove `json:"moves"`
	Sprites struct {
		FrontDefault     *string `json:"front_default"`
		FrontFemale      string  `json:"front_female"`
		FrontShiny       string  `json:"front_shiny"`
		FrontShinyFemale string  `json:"front_shiny_female"`
		BackDefault      string  `json:"back_default"`
		BackFemale       string  `json:"back_female"`
		BackShiny        string  `json:"back_shiny"`
		BackShinyFemale  string  `json:"back_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  string `json:"front_female"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
			} `json:"official_artwork"`
		} `json:"other"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int             `json:"base_stat"`
		Effort   int             `json:"effort"`
		Stat     ResourcePointer `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int             `json:"slot"`
		Type ResourcePointer `json:"type"`
	} `json:"types"`
}

type PokemonMove struct {
	Move                ResourcePointer `json:"move"`
	VersionGroupDetails []struct {
		LevelLearnedAt  int             `json:"level_learned_at"`
		VersionGroup    ResourcePointer `json:"version_group"`
		MoveLearnMethod ResourcePointer `json:"move_learn_method"`
	} `json:"version_group_details"`
}

type PokemonSpecies struct {
	ID                   int                `json:"id"`
	Name                 string             `json:"name"`
	Order                int                `json:"order"`
	GenderRate           int                `json:"gender_rate"`
	CaptureRate          int                `json:"capture_rate"`
	BaseHappiness        int                `json:"base_happiness"`
	IsBaby               bool               `json:"is_baby"`
	IsLegendary          bool               `json:"is_legendary"`
	IsMythical           bool               `json:"is_mythical"`
	HatchCounter         int                `json:"hatch_counter"`
	HasGenderDifferences bool               `json:"has_gender_differences"`
	FormsSwitchable      bool               `json:"forms_switchable"`
	GrowthRate           ResourcePointer    `json:"growth_rate"`
	EggGroups            []ResourcePointer  `json:"egg_groups"`
	Color                ResourcePointer    `json:"color"`
	Shape                ResourcePointer    `json:"shape"`
	EvolvesFromSpecies   ResourcePointer    `json:"evolves_from_species"`
	EvolutionChain       ResourcePointer    `json:"evolution_chain"`
	Habitat              *ResourcePointer   `json:"habitat"`
	Names                []TranslatedName   `json:"names"`
	FlavourTextEntries   []FlavourTextEntry `json:"flavor_text_entries"`
	Genera               []Genus            `json:"genera"`
	PokedexNumbers       []struct {
		EntryNumber int             `json:"entry_number"`
		Pokedex     ResourcePointer `json:"pokedex"`
	} `json:"pokedex_numbers"`
	Varieties []struct {
		IsDefault bool            `json:"is_default"`
		Pokemon   ResourcePointer `json:"pokemon"`
	} `json:"varieties"`
}

type Region struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Locations      []ResourcePointer `json:"locations"`
	MainGeneration ResourcePointer   `json:"main_generation"`
	Names          []TranslatedName  `json:"names"`
	Pokedexes      []ResourcePointer `json:"pokedexes"`
	VersionGroups  []ResourcePointer `json:"version_groups"`
}

type ResourcePointerList struct {
	Count   int               `json:"count"`
	Results []ResourcePointer `json:"results"`
}

type ResourcePointer struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stat struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`
	GameIndex    int              `json:"game_index"`
	IsBattleOnly bool             `json:"is_battle_only"`
	Names        []TranslatedName `json:"names"`
}

type Target struct {
	Id    int              `json:"id"`
	Name  string           `json:"name"`
	Names []TranslatedName `json:"names"`
}

type TranslatedName struct {
	Name     string          `json:"name"`
	Language ResourcePointer `json:"language"`
}

type Type struct {
	Id              int               `json:"id"`
	Name            string            `json:"name"`
	Names           []TranslatedName  `json:"names"`
	Moves           []ResourcePointer `json:"moves"`
	DamageRelations struct {
		NoDamageTo       []ResourcePointer `json:"no_damage_to"`
		HalfDamageTo     []ResourcePointer `json:"half_damage_to"`
		DoubleDamageTo   []ResourcePointer `json:"double_damage_to"`
		NoDamageFrom     []ResourcePointer `json:"no_damage_from"`
		HalfDamageFrom   []ResourcePointer `json:"half_damage_from"`
		DoubleDamageFrom []ResourcePointer `json:"double_damage_from"`
	} `json:"damage_relations"`
	Pokemon []struct {
		Slot    int             `json:"slot"`
		Pokemon ResourcePointer `json:"pokemon"`
	} `json:"pokemon"`
}

type VersionGroup struct {
	Id               int               `json:"id"`
	Name             string            `json:"name"`
	Order            int               `json:"order"`
	Generation       ResourcePointer   `json:"generation"`
	MoveLearnMethods []ResourcePointer `json:"move_learn_methods"`
	Pokedexes        []ResourcePointer `json:"pokedexes"`
	Regions          []ResourcePointer `json:"regions"`
	Versions         []ResourcePointer `json:"versions"`
}
