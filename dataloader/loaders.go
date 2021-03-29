package dataloader

import (
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type key string

const loadersKey key = "dataloaders"

type Loaders struct {
	MovesByPokemonId MoveListLoader
	PokemonByMoveId  PokemonListLoader
	PokemonByTypeId  PokemonListLoader
	TypesByPokemonId TypeListLoader
	TypeByTypeId     TypeLoader
}

func Middleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
				MovesByPokemonId: MoveListLoader{
					maxBatch: 1000,
					wait:     1 * time.Millisecond,
					fetch: func(ids []string) ([]*model.MoveList, []error) {
						movesByPokemonId := map[string]*model.MoveList{}
						placeholders := make([]string, len(ids))
						args := make([]interface{}, len(ids))
						for i := 0; i < len(ids); i++ {
							placeholders[i] = fmt.Sprintf("$%d", i+1)
							args[i] = ids[i]
						}

						query := "SELECT id, name, slug, accuracy, pp, power, damage_class_enum, effect, effect_chance, target, type_id, pokemon_move.pokemon_id FROM moves LEFT JOIN pokemon_move ON moves.id = pokemon_move.move_id WHERE pokemon_move.pokemon_id IN (" + strings.Join(placeholders, ",") + ")"

						log.Logger.WithField("args", args).Debug(query)
						rows, err := db.QueryContext(r.Context(),
							query,
							args...,
						)
						if err != nil {
							panic(fmt.Errorf("error fetching moves for pokemon: %w", err))
						}

						defer rows.Close()
						for rows.Next() {
							var m model.Move
							var pokemonId string
							err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Accuracy, &m.PP, &m.Power, &m.DamageClass, &m.Effect, &m.EffectChance, &m.Target, &m.TypeId, &pokemonId)
							if err != nil {
								panic(fmt.Errorf("error scanning result in MovesByPokemonId: %w", err))
							}

							_, ok := movesByPokemonId[pokemonId]
							if !ok {
								ml := model.NewEmptyMoveList()
								movesByPokemonId[pokemonId] = &ml
							}

							movesByPokemonId[pokemonId].AddMove(&m)
						}

						moveList := make([]*model.MoveList, len(ids))
						for i, id := range ids {
							moveList[i] = movesByPokemonId[id]
							i++
						}

						return moveList, nil
					},
				},
				PokemonByMoveId: PokemonListLoader{
					maxBatch: 1000,
					wait:     1 * time.Millisecond,
					fetch: func(ids []string) ([]*model.PokemonList, []error) {
						pokemonByMoveId := map[string]*model.PokemonList{}
						placeholders := make([]string, len(ids))
						args := make([]interface{}, len(ids))
						for i := 0; i < len(ids); i++ {
							placeholders[i] = fmt.Sprintf("$%d", i+1)
							args[i] = ids[i]
						}

						query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, pokemon_move.move_id FROM pokemon LEFT JOIN pokemon_move ON pokemon.id = pokemon_move.pokemon_id WHERE pokemon_move.move_id IN (" + strings.Join(placeholders, ",") + ")"

						log.Logger.WithField("args", args).Debug(query)
						rows, err := db.QueryContext(r.Context(),
							query,
							args...,
						)
						if err != nil {
							panic(fmt.Errorf("error fetching pokemon for move: %w", err))
						}

						defer rows.Close()
						for rows.Next() {
							var pkmn model.Pokemon
							var moveId string
							err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &moveId)
							if err != nil {
								panic(fmt.Errorf("error scanning result in PokemonByMoveId: %w", err))
							}

							_, ok := pokemonByMoveId[moveId]
							if !ok {
								pl := model.NewEmptyPokemonList()
								pokemonByMoveId[moveId] = &pl
							}

							pokemonByMoveId[moveId].AddPokemon(&pkmn)
						}

						pokemonList := make([]*model.PokemonList, len(ids))
						for i, id := range ids {
							pokemonList[i] = pokemonByMoveId[id]
							i++
						}

						return pokemonList, nil
					},
				},
				PokemonByTypeId: PokemonListLoader{
					maxBatch: 1000,
					wait:     1 * time.Millisecond,
					fetch: func(ids []string) ([]*model.PokemonList, []error) {
						pokemonByTypeId := map[string]*model.PokemonList{}
						placeholders := make([]string, len(ids))
						args := make([]interface{}, len(ids))
						for i := 0; i < len(ids); i++ {
							placeholders[i] = fmt.Sprintf("$%d", i+1)
							args[i] = ids[i]
						}

						query := "SELECT id, name, slug, pokedex_id, sprite, hp, attack, defense, special_attack, special_defense, speed, is_baby, is_legendary, is_mythical, description, pokemon_type.type_id FROM pokemon LEFT JOIN pokemon_type ON pokemon.id = pokemon_type.pokemon_id WHERE pokemon_type.type_id IN (" + strings.Join(placeholders, ",") + ")"

						log.Logger.WithField("args", args).Debug(query)
						rows, err := db.QueryContext(r.Context(),
							query,
							args...,
						)
						if err != nil {
							panic(fmt.Errorf("error fetching pokemon for type: %w", err))
						}

						defer rows.Close()
						for rows.Next() {
							var pkmn model.Pokemon
							var typeId string
							err := rows.Scan(&pkmn.ID, &pkmn.Name, &pkmn.Slug, &pkmn.PokedexId, &pkmn.Sprite, &pkmn.HP, &pkmn.Attack, &pkmn.Defense, &pkmn.SpecialAttack, &pkmn.SpecialDefense, &pkmn.Speed, &pkmn.IsBaby, &pkmn.IsLegendary, &pkmn.IsMythical, &pkmn.Description, &typeId)
							if err != nil {
								panic(fmt.Errorf("error scanning result in PokemonByTypeId: %w", err))
							}

							_, ok := pokemonByTypeId[typeId]
							if !ok {
								pl := model.NewEmptyPokemonList()
								pokemonByTypeId[typeId] = &pl
							}

							pokemonByTypeId[typeId].AddPokemon(&pkmn)
						}

						pokemonList := make([]*model.PokemonList, len(ids))
						for i, id := range ids {
							pokemonList[i] = pokemonByTypeId[id]
							i++
						}

						return pokemonList, nil
					},
				},
				TypesByPokemonId: TypeListLoader{
					maxBatch: 1000,
					wait:     1 * time.Millisecond,
					fetch: func(ids []string) ([]*model.TypeList, []error) {
						typesByPokemonId := map[string]*model.TypeList{}
						placeholders := make([]string, len(ids))
						args := make([]interface{}, len(ids))
						for i := 0; i < len(ids); i++ {
							placeholders[i] = fmt.Sprintf("$%d", i+1)
							args[i] = ids[i]
						}

						query := "SELECT id, name, slug, pokemon_type.pokemon_id FROM types LEFT JOIN pokemon_type ON types.id = pokemon_type.type_id WHERE pokemon_type.pokemon_id IN (" + strings.Join(placeholders, ",") + ")"

						log.Logger.WithField("args", args).Debug(query)
						rows, err := db.QueryContext(r.Context(),
							query,
							args...,
						)
						if err != nil {
							panic(fmt.Errorf("error fetching types for pokemon: %w", err))
						}

						defer rows.Close()
						for rows.Next() {
							var t model.Type
							var pokemonId string
							err := rows.Scan(&t.ID, &t.Name, &t.Slug, &pokemonId)
							if err != nil {
								panic(fmt.Errorf("error scanning result in TypesByPokemonId: %w", err))
							}

							_, ok := typesByPokemonId[pokemonId]
							if !ok {
								tl := model.NewEmptyTypeList()
								typesByPokemonId[pokemonId] = &tl
							}

							typesByPokemonId[pokemonId].AddType(&t)
						}

						typeList := make([]*model.TypeList, len(ids))
						for i, id := range ids {
							typeList[i] = typesByPokemonId[id]
							i++
						}

						return typeList, nil
					},
				},
				TypeByTypeId: TypeLoader{
					maxBatch: 1000,
					wait:     1 * time.Millisecond,
					fetch: func(ids []string) ([]*model.Type, []error) {
						typesByTypeId := map[string]*model.Type{}
						placeholders := make([]string, len(ids))
						args := make([]interface{}, len(ids))
						for i := 0; i < len(ids); i++ {
							placeholders[i] = fmt.Sprintf("$%d", i+1)
							args[i] = ids[i]
						}

						query := "SELECT id, name, slug FROM types WHERE id IN (" + strings.Join(placeholders, ",") + ")"

						log.Logger.WithField("args", args).Debug(query)
						rows, err := db.QueryContext(r.Context(),
							query,
							args...,
						)
						if err != nil {
							panic(fmt.Errorf("error fetching types: %w", err))
						}

						defer rows.Close()
						for rows.Next() {
							var t model.Type
							err := rows.Scan(&t.ID, &t.Name, &t.Slug)
							if err != nil {
								panic(fmt.Errorf("error scanning result in TypeByTypeId: %w", err))
							}

							typesByTypeId[t.ID] = &t
						}

						types := make([]*model.Type, len(ids))
						for i, id := range ids {
							types[i] = typesByTypeId[id]
							i++
						}

						return types, nil
					},
				},
			})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
