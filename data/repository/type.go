package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"fmt"
)

type Type struct {
	client *db.PrismaClient
}

func NewType(client *db.PrismaClient) Type {
	return Type{
		client: client,
	}
}

func (r Type) GetTypes(ctx context.Context) (*model.TypeList, error) {
	types := model.NewEmptyTypeList()

	results, err := r.client.Type.FindMany().Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &types, nil
	}

	if err != nil {
		return &types, fmt.Errorf("error getting types: %s", err)
	}

	for _, result := range results {
		t := model.NewTypeFromDb(result)
		types.AddType(&t)
	}

	return &types, nil
}

func (r Type) GetTypeById(ctx context.Context, id string) (*model.Type, error) {
	result, err := r.client.Type.FindUnique(db.Type.ID.Equals(id)).Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find type by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error getting type by id: %s, error: %s", id, err)
	}

	t := model.NewTypeFromDb(*result)
	return &t, nil
}

func (r Type) TypeByIdWithDamageRelationDataLoader(ctx context.Context, damageRelation db.DamageRelation) func(ids []string) ([]*model.TypeList, []error) {
	return func(ids []string) ([]*model.TypeList, []error) {
		typeListsById := map[string]*model.TypeList{}
		results, err := r.client.TypeDamageRelation.
			FindMany(db.TypeDamageRelation.TypeAID.In(ids), db.TypeDamageRelation.DamageRelation.Equals(damageRelation)).
			With(db.TypeDamageRelation.TypeB.Fetch()).
			Exec(ctx)
		typeLists := make([]*model.TypeList, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				tl := model.NewEmptyTypeList()
				typeLists[i] = &tl
				errors[i] = fmt.Errorf("error loading type by damage relation %s & id %s in dataloader %w", damageRelation, id, err)
			}

			return typeLists, errors
		}

		if len(results) == 0 {
			for i := range ids {
				tl := model.NewEmptyTypeList()
				typeLists[i] = &tl
			}

			return typeLists, nil
		}

		for _, result := range results {
			tl := typeListsById[result.TypeAID]
			if tl == nil {
				empty := model.NewEmptyTypeList()
				typeListsById[result.TypeAID] = &empty
			}
			t := model.NewTypeFromDb(*result.TypeB())
			typeListsById[result.TypeAID].AddType(&t)
		}

		for i, id := range ids {
			typeList := typeListsById[id]

			if typeList == nil {
				empty := model.NewEmptyTypeList()
				typeList = &empty
			}

			typeLists[i] = typeList
		}

		return typeLists, nil
	}
}

func (r Type) TypeByIdDataLoader(ctx context.Context) func(ids []string) ([]*model.Type, []error) {
	return func(ids []string) ([]*model.Type, []error) {
		typesById := map[string]*model.Type{}
		results, err := r.client.Type.FindMany(db.Type.ID.In(ids)).Exec(ctx)
		types := make([]*model.Type, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				errors[i] = fmt.Errorf("error loading type by id %s in dataloader %w", id, err)
			}

			return types, errors
		}

		for _, result := range results {
			t := model.NewTypeFromDb(result)
			typesById[t.ID] = &t
		}

		for i, id := range ids {
			types[i] = typesById[id]
		}

		return types, nil
	}
}
