package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"fmt"
)

type Team struct {
	client *db.PrismaClient
}

func NewTeam(client *db.PrismaClient) Team {
	return Team{
		client: client,
	}
}

func (r Team) GetTeams(ctx context.Context) (*model.TeamList, error) {
	teams := model.NewEmptyTeamList()

	results, err := r.client.Team.FindMany().
		With(db.Team.TeamMembers.Fetch().With(db.TeamMember.Moves.Fetch())).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &teams, nil
	}

	if err != nil {
		return &teams, fmt.Errorf("error getting teams: %s", err)
	}

	for _, result := range results {
		team := model.NewTeamFromDb(result)
		teams.AddTeam(&team)
	}

	return &teams, nil
}

func (r Team) GetTeamById(ctx context.Context, id string) (*model.Team, error) {
	result, err := r.client.Team.
		FindUnique(db.Team.ID.Equals(id)).
		With(db.Team.TeamMembers.Fetch().With(db.TeamMember.Moves.Fetch())).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find team by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error getting team by id: %s, error: %s", id, err)
	}

	team := model.NewTeamFromDb(*result)
	return &team, nil
}
