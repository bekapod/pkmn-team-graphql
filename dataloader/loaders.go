//go:generate go run github.com/vektah/dataloaden PokemonTypeLoader string "*bekapod/pkmn-team-graphql/data/model.TypeList"

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
	TypesByPokemonId PokemonTypeLoader
}

func Middleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
				TypesByPokemonId: PokemonTypeLoader{
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
								panic(fmt.Errorf("error scanning result in GetAllTypes: %w", err))
							}

							_, ok := typesByPokemonId[pokemonId]
							if !ok {
								typeList := model.NewEmptyTypeList()
								typesByPokemonId[pokemonId] = &typeList
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
			})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
