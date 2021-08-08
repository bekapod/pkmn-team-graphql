package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Type struct {
	client *db.PrismaClient
}

func NewType(client *db.PrismaClient) Type {
	return Type{
		client: client,
	}
}

func (r Type) GetTypes(ctx context.Context) (*model.TypeConnection, error) {
	types := model.NewEmptyTypeConnection()

	results, err := r.client.Type.FindMany().Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &types, nil
	}

	if err != nil {
		log.Logger.WithError(err).Error("error getting types")
		return &types, fmt.Errorf("error getting types")
	}

	for _, result := range results {
		t := model.NewTypeEdgeFromDb(result)
		types.AddEdge(&t)
	}

	return &types, nil
}

func (r Type) GetTypeById(ctx context.Context, id string) (*model.Type, error) {
	result, err := r.client.Type.FindUnique(db.Type.ID.Equals(id)).Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		log.Logger.WithField("id", id).WithContext(ctx).Info("couldn't find type by id")
		return nil, fmt.Errorf("couldn't find type by id: %s", id)
	}

	if err != nil {
		log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error getting type by id")
		return nil, fmt.Errorf("error getting type by id: %s", id)
	}

	t := model.NewTypeFromDb(*result)
	return &t, nil
}

func (r Type) TypeByIdWithDamageRelationDataLoader(ctx context.Context, damageRelation db.DamageRelation) func(ids []string) ([]*model.TypeConnection, []error) {
	return func(ids []string) ([]*model.TypeConnection, []error) {
		typeConnectionsById := map[string]*model.TypeConnection{}
		results, err := r.client.TypeDamageRelation.
			FindMany(db.TypeDamageRelation.TypeAID.In(ids), db.TypeDamageRelation.DamageRelation.Equals(damageRelation)).
			With(db.TypeDamageRelation.TypeB.Fetch()).
			Exec(ctx)
		typeConnections := make([]*model.TypeConnection, len(ids))

		if err != nil {
			errors := make([]error, len(ids))
			for i, id := range ids {
				tl := model.NewEmptyTypeConnection()
				typeConnections[i] = &tl
				log.Logger.WithFields(logrus.Fields{"id": id, "damageRelation": damageRelation}).WithError(err).WithContext(ctx).Error("error loading type by damage relation")
				errors[i] = fmt.Errorf("error loading type by damage relation %s & id %s", damageRelation, id)
			}

			return typeConnections, errors
		}

		if len(results) == 0 {
			for i := range ids {
				tl := model.NewEmptyTypeConnection()
				typeConnections[i] = &tl
			}

			return typeConnections, nil
		}

		for _, result := range results {
			tl := typeConnectionsById[result.TypeAID]
			if tl == nil {
				empty := model.NewEmptyTypeConnection()
				typeConnectionsById[result.TypeAID] = &empty
			}
			t := model.NewTypeEdgeFromDb(*result.TypeB())
			typeConnectionsById[result.TypeAID].AddEdge(&t)
		}

		for i, id := range ids {
			typeConnection := typeConnectionsById[id]

			if typeConnection == nil {
				empty := model.NewEmptyTypeConnection()
				typeConnection = &empty
			}

			typeConnections[i] = typeConnection
		}

		return typeConnections, nil
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
				log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error loading type by id")
				errors[i] = fmt.Errorf("error loading type by id %s", id)
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
