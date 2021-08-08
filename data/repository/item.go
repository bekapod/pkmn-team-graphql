package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"fmt"
)

type Item struct {
	client *db.PrismaClient
}

func NewItem(client *db.PrismaClient) Item {
	return Item{
		client: client,
	}
}

func (r Item) ItemByIdDataLoader(ctx context.Context) func(ids []string) ([]*model.Item, []error) {
	return func(ids []string) ([]*model.Item, []error) {
		itemsById := map[string]*model.Item{}
		results, err := r.client.Item.FindMany(db.Item.ID.In(ids)).Exec(ctx)
		items := make([]*model.Item, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading item by id")
				errors[i] = fmt.Errorf("error loading item by id %s", id)
			}

			return items, errors
		}

		for _, result := range results {
			i := model.NewItemFromDb(result)
			itemsById[i.ID] = &i
		}

		for i, id := range ids {
			items[i] = itemsById[id]
		}

		return items, nil
	}
}
