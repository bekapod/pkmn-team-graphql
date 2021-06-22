package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"context"
	"errors"
	"fmt"

	"github.com/prisma/prisma-client-go/runtime/transaction"
)

type Team struct {
	client *db.PrismaClient
}

func NewTeam(client *db.PrismaClient) Team {
	return Team{
		client: client,
	}
}

func (r Team) GetTeams(ctx context.Context) (*model.TeamConnection, error) {
	teams := model.NewEmptyTeamConnection()

	results, err := r.client.Team.FindMany().
		With(db.Team.TeamMembers.Fetch().With(
			db.TeamMember.Moves.Fetch().OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).OrderBy(db.TeamMember.Slot.Order(db.ASC))).
		OrderBy(db.Team.CreatedAt.Order(db.DESC)).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &teams, nil
	}

	if err != nil {
		return &teams, fmt.Errorf("error getting teams: %s", err)
	}

	for _, result := range results {
		team := model.NewTeamEdgeFromDb(result)
		teams.AddEdge(&team)
	}

	return &teams, nil
}

func (r Team) GetTeamById(ctx context.Context, id string) (*model.Team, error) {
	result, err := r.client.Team.
		FindUnique(db.Team.ID.Equals(id)).
		With(db.Team.TeamMembers.Fetch().With(
			db.TeamMember.Moves.Fetch().OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).OrderBy(db.TeamMember.Slot.Order(db.ASC))).
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

func (r Team) CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error) {
	result, err := r.client.Team.
		CreateOne(db.Team.Name.Set(input.Name)).
		Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("error creating team %s", err)
	}

	transactions := make([]transaction.Param, 0)
	for _, teamMemberInput := range input.Members {
		tx := r.client.TeamMember.
			CreateOne(
				db.TeamMember.Slot.Set(teamMemberInput.Slot),
				db.TeamMember.Pokemon.Link(db.Pokemon.ID.Equals(teamMemberInput.PokemonID)),
				db.TeamMember.Team.Link(db.Team.ID.Equals(result.ID)),
			).
			Tx()

		transactions = append(transactions, tx)
	}

	err2 := r.client.Prisma.Transaction(transactions...).Exec(ctx)
	if err2 != nil {
		r.client.Team.FindUnique(db.Team.ID.Equals(result.ID)).Delete().Exec(ctx)
		return nil, fmt.Errorf("error creating team members %s", err2)
	}

	return r.GetTeamById(ctx, result.ID)
}

func (r Team) UpdateTeam(ctx context.Context, input model.UpdateTeamInput) (*model.Team, error) {
	result, err := r.client.Team.
		FindUnique(db.Team.ID.Equals(input.ID)).
		Update(
			db.Team.Name.SetIfPresent(input.Name),
		).
		Exec(ctx)

	if err != nil {
		return nil, fmt.Errorf("error updating team %s", err)
	}

	transactions := make([]transaction.Param, 0)
	for _, teamMemberInput := range input.Members {
		var tx transaction.Param
		if teamMemberInput.ID != nil {
			tx = r.client.TeamMember.
				FindUnique(db.TeamMember.ID.Equals(*teamMemberInput.ID)).
				Update(
					db.TeamMember.Slot.SetIfPresent(teamMemberInput.Slot),
					db.TeamMember.Pokemon.Link(db.Pokemon.ID.EqualsIfPresent(teamMemberInput.PokemonID)),
				).
				Tx()
		} else {
			tx = r.client.TeamMember.
				CreateOne(
					db.TeamMember.Slot.SetIfPresent(teamMemberInput.Slot),
					db.TeamMember.Pokemon.Link(db.Pokemon.ID.EqualsIfPresent(teamMemberInput.PokemonID)),
					db.TeamMember.Team.Link(db.Team.ID.Equals(result.ID)),
				).
				Tx()
		}

		transactions = append(transactions, tx)
	}

	err2 := r.client.Prisma.Transaction(transactions...).Exec(ctx)
	if err2 != nil {
		r.client.Team.FindUnique(db.Team.ID.Equals(result.ID)).Delete().Exec(ctx)
		return nil, fmt.Errorf("error updating team members %s", err2)
	}

	return r.GetTeamById(ctx, result.ID)
}

func (r Team) DeleteTeam(ctx context.Context, id string) (*model.Team, error) {
	result, err := r.client.Team.
		FindUnique(db.Team.ID.Equals(id)).
		With(db.Team.TeamMembers.Fetch().With(
			db.TeamMember.Moves.Fetch().OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).OrderBy(db.TeamMember.Slot.Order(db.ASC))).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find team by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error finding team by id %s: %s", id, err)
	}

	_, err2 := r.client.TeamMember.FindMany(db.TeamMember.TeamID.Equals(id)).Delete().Exec(ctx)
	if err2 != nil {
		return nil, fmt.Errorf("error deleting team members by team id %s: %s", id, err2)
	}

	_, err3 := r.client.Team.FindUnique(db.Team.ID.Equals(id)).Delete().Exec(ctx)
	if err3 != nil {
		return nil, fmt.Errorf("error deleting team by id %s: %s", id, err3)
	}

	team := model.NewTeamFromDb(*result)
	return &team, nil
}

func (r Team) DeleteTeamMember(ctx context.Context, id string) (*model.TeamMember, error) {
	result, err := r.client.TeamMember.
		FindUnique(db.TeamMember.ID.Equals(id)).
		With(
			db.TeamMember.Team.Fetch().With(
				db.Team.TeamMembers.Fetch().With(
					db.TeamMember.Moves.Fetch().OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
				).OrderBy(db.TeamMember.Slot.Order(db.ASC)),
			),
			db.TeamMember.Moves.Fetch().OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).
		Delete().
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return nil, fmt.Errorf("couldn't find team member by id: %s", id)
	}

	if err != nil {
		return nil, fmt.Errorf("error deleting team member by id %s: %s", id, err)
	}

	teamMember := model.NewTeamMemberFromDb(*result)
	return &teamMember, nil
}
