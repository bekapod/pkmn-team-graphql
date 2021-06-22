package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"errors"
	"fmt"
)

type Ability struct {
	client *db.PrismaClient
}

func NewAbility(client *db.PrismaClient) Ability {
	return Ability{
		client: client,
	}
}

func (r Ability) GetAbilities(ctx context.Context) (*model.AbilityConnection, error) {
	abilities := model.NewEmptyAbilityConnection()

	results, err := r.client.Ability.FindMany().Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &abilities, nil
	}

	if err != nil {
		log.Logger.WithField("r", results).Debug("resultssssss")
		return &abilities, fmt.Errorf("error getting abilities: %s", err)
	}

	for _, result := range results {
		a := model.NewAbilityEdgeFromDb(result)
		abilities.AddEdge(&a)
	}

	return &abilities, nil
}

func (r Ability) GetAbilityById(ctx context.Context, id string) (*model.Ability, error) {
	result, err := r.client.Ability.FindUnique(db.Ability.ID.Equals(id)).Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find ability by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error getting ability by id: %s, error: %s", id, err)
	}

	a := model.NewAbilityFromDb(*result)
	return &a, nil
}

func (r Ability) AbilityByIdDataLoader(ctx context.Context) func(ids []string) ([]*model.Ability, []error) {
	return func(ids []string) ([]*model.Ability, []error) {
		abilitiesById := map[string]*model.Ability{}
		results, err := r.client.Ability.FindMany(db.Ability.ID.In(ids)).Exec(ctx)
		abilities := make([]*model.Ability, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				errors[i] = fmt.Errorf("error loading ability by id %s in dataloader %w", id, err)
			}

			return abilities, errors
		}

		for _, result := range results {
			a := model.NewAbilityFromDb(result)
			abilitiesById[a.ID] = &a
		}

		for i, id := range ids {
			abilities[i] = abilitiesById[id]
		}

		return abilities, nil
	}
}
