package repository

import (
	"bekapod/pkmn-team-graphql/data/db"
	"bekapod/pkmn-team-graphql/data/model"
	"bekapod/pkmn-team-graphql/log"
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
			db.TeamMember.Moves.Fetch().With(
				db.TeamMemberMove.PokemonMove.Fetch(),
			).OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).OrderBy(db.TeamMember.Slot.Order(db.ASC))).
		OrderBy(db.Team.CreatedAt.Order(db.DESC)).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		return &teams, nil
	}

	if err != nil {
		log.Logger.WithError(err).WithContext(ctx).Error("error getting teams")
		return &teams, fmt.Errorf("error getting teams")
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
			db.TeamMember.Moves.Fetch().With(db.TeamMemberMove.PokemonMove.Fetch()).OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).OrderBy(db.TeamMember.Slot.Order(db.ASC))).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		log.Logger.WithField("id", id).WithContext(ctx).Info("couldn't find team by id")
		return nil, fmt.Errorf("couldn't find team by id: %s", id)
	}

	if err != nil {
		log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error getting team by id")
		return nil, fmt.Errorf("error getting team by id")
	}

	team := model.NewTeamFromDb(*result)
	return &team, nil
}

func (r Team) GetTeamMemberById(ctx context.Context, id string) (*model.TeamMember, error) {
	result, err := r.client.TeamMember.
		FindUnique(db.TeamMember.ID.Equals(id)).
		With(db.TeamMember.Moves.Fetch().With(db.TeamMemberMove.PokemonMove.Fetch()).OrderBy(db.TeamMemberMove.Slot.Order(db.ASC))).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		log.Logger.WithField("id", id).WithContext(ctx).Info("couldn't get team member by id")
		return nil, fmt.Errorf("couldn't get team member by id: %s", id)
	}

	if err != nil {
		log.Logger.WithField("id", id).WithError(err).WithContext(ctx).Error("error getting team member by id")
		return nil, fmt.Errorf("error getting team member by id: %s", id)
	}

	teamMember := model.NewTeamMemberFromDb(*result)
	return &teamMember, nil
}

func (r Team) CreateTeam(ctx context.Context, input model.CreateTeamInput) (*model.Team, error) {
	result, err := r.client.Team.
		CreateOne(db.Team.Name.Set(input.Name)).
		Exec(ctx)

	if err != nil {
		log.Logger.WithError(err).WithContext(ctx).Error("error creating team")
		return nil, fmt.Errorf("error creating team")
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
		log.Logger.WithError(err2).WithContext(ctx).Error("error creating team members")
		return nil, fmt.Errorf("error creating team members")
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
		log.Logger.WithError(err).WithContext(ctx).Error("error updating team")
		return nil, fmt.Errorf("error updating team")
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
		log.Logger.WithError(err2).WithContext(ctx).Error("error updating team members")
		return nil, fmt.Errorf("error updating team members")
	}

	return r.GetTeamById(ctx, result.ID)
}

func (r Team) UpdateTeamMember(ctx context.Context, input model.UpdateTeamMemberInput) (*model.TeamMember, error) {
	result, err := r.client.TeamMember.
		FindUnique(db.TeamMember.ID.Equals(input.ID)).
		Update(
			db.TeamMember.Pokemon.Link(db.Pokemon.ID.EqualsIfPresent(input.PokemonID)),
			db.TeamMember.Slot.SetIfPresent(input.Slot),
		).
		Exec(ctx)

	if err != nil {
		log.Logger.WithField("id", input.ID).WithContext(ctx).WithError(err).Error("error updating team member")
		return nil, fmt.Errorf("error updating team member %s", input.ID)
	}

	transactions := make([]transaction.Param, 0)
	for _, teamMemberMoveInput := range input.Moves {
		var tx transaction.Param
		if teamMemberMoveInput.ID != nil {
			tx = r.client.TeamMemberMove.
				FindUnique(db.TeamMemberMove.ID.Equals(*teamMemberMoveInput.ID)).
				Update(
					db.TeamMemberMove.Slot.SetIfPresent(teamMemberMoveInput.Slot),
					db.TeamMemberMove.PokemonMove.Link(db.PokemonMove.ID.EqualsIfPresent(teamMemberMoveInput.PokemonMoveID)),
				).
				Tx()
		} else {
			tx = r.client.TeamMemberMove.
				CreateOne(
					db.TeamMemberMove.Slot.SetIfPresent(teamMemberMoveInput.Slot),
					db.TeamMemberMove.TeamMember.Link(db.TeamMember.ID.Equals(input.ID)),
					db.TeamMemberMove.PokemonMove.Link(db.PokemonMove.ID.EqualsIfPresent(teamMemberMoveInput.PokemonMoveID)),
				).
				Tx()
		}

		transactions = append(transactions, tx)
	}

	err2 := r.client.Prisma.Transaction(transactions...).Exec(ctx)
	if err2 != nil {
		log.Logger.WithField("id", result.ID).WithError(err2).WithContext(ctx).Error("error updating team member moves")
		return nil, fmt.Errorf("error updating team member moves")
	}

	return r.GetTeamMemberById(ctx, result.ID)
}

func (r Team) DeleteTeams(ctx context.Context, ids []string) ([]*model.Team, error) {
	results, err := r.client.Team.
		FindMany(db.Team.ID.In(ids)).
		With(db.Team.TeamMembers.Fetch().With(
			db.TeamMember.Moves.Fetch().With(db.TeamMemberMove.PokemonMove.Fetch()).OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).OrderBy(db.TeamMember.Slot.Order(db.ASC))).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		log.Logger.WithField("ids", ids).WithContext(ctx).Info("couldn't find teams with ids for deletion")
		return nil, fmt.Errorf("couldn't find team by id: %s", ids)
	}

	if err != nil {
		log.Logger.WithField("ids", ids).WithError(err).WithContext(ctx).Error("error finding teams by ids for deletion")
		return nil, fmt.Errorf("error finding teams by ids %s", ids)
	}

	_, err2 := r.client.TeamMemberMove.FindMany(db.TeamMemberMove.TeamMember.Where(db.TeamMember.TeamID.In(ids))).Delete().Exec(ctx)
	if err2 != nil {
		log.Logger.WithField("ids", ids).WithError(err2).WithContext(ctx).Error("error deleting team member moves by team ids")
		return nil, fmt.Errorf("error deleting team member moves by team ids %s", ids)
	}

	_, err3 := r.client.TeamMember.FindMany(db.TeamMember.TeamID.In(ids)).Delete().Exec(ctx)
	if err3 != nil {
		log.Logger.WithField("ids", ids).WithError(err3).WithContext(ctx).Error("error deleting team members by team ids")
		return nil, fmt.Errorf("error deleting team members by team ids %s", ids)
	}

	_, err4 := r.client.Team.FindMany(db.Team.ID.In(ids)).Delete().Exec(ctx)
	if err4 != nil {
		log.Logger.WithField("ids", ids).WithError(err4).WithContext(ctx).Error("error deleting team by id")
		return nil, fmt.Errorf("error deleting team by id %s", ids)
	}

	teams := make([]*model.Team, 0)
	for _, result := range results {
		team := model.NewTeamFromDb(result)
		teams = append(teams, &team)
	}
	return teams, nil
}

func (r Team) DeleteTeamMembers(ctx context.Context, ids []string) ([]*model.TeamMember, error) {
	results, err := r.client.TeamMember.
		FindMany(db.TeamMember.ID.In(ids)).
		With(
			db.TeamMember.Team.Fetch().With(
				db.Team.TeamMembers.Fetch().With(
					db.TeamMember.Moves.Fetch().With(db.TeamMemberMove.PokemonMove.Fetch()).OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
				).OrderBy(db.TeamMember.Slot.Order(db.ASC)),
			),
			db.TeamMember.Moves.Fetch().With(db.TeamMemberMove.PokemonMove.Fetch()).OrderBy(db.TeamMemberMove.Slot.Order(db.ASC)),
		).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		log.Logger.WithField("ids", ids).WithContext(ctx).Info("couldn't find team members by ids for deletion")
		return nil, fmt.Errorf("couldn't find team members by ids for deletion: %s", ids)
	}

	if err != nil {
		log.Logger.WithField("ids", ids).WithError(err).WithContext(ctx).Error("error getting team members by ids")
		return nil, fmt.Errorf("error getting team members by ids %s", ids)
	}

	_, err2 := r.client.TeamMemberMove.FindMany(db.TeamMemberMove.TeamMemberID.In(ids)).Delete().Exec(ctx)

	if err2 != nil {
		log.Logger.WithField("ids", ids).WithError(err2).WithContext(ctx).Error("error deleting moves for team members by ids")
		return nil, fmt.Errorf("error deleting moves for team members by ids %s", ids)
	}

	_, err3 := r.client.TeamMember.FindMany(db.TeamMember.ID.In(ids)).Delete().Exec(ctx)

	if err3 != nil {
		log.Logger.WithField("ids", ids).WithError(err3).WithContext(ctx).Error("error deleting team members by ids")
		return nil, fmt.Errorf("error deleting team members by ids %s", ids)
	}

	teamMembers := make([]*model.TeamMember, 0)
	for _, result := range results {
		teamMember := model.NewTeamMemberFromDb(result)
		teamMembers = append(teamMembers, &teamMember)
	}
	return teamMembers, nil
}

func (r Team) DeleteTeamMemberMoves(ctx context.Context, ids []string) ([]*model.Move, error) {
	results, err := r.client.TeamMemberMove.
		FindMany(db.TeamMemberMove.ID.In(ids)).
		With(db.TeamMemberMove.PokemonMove.Fetch().With(db.PokemonMove.Move.Fetch())).
		Exec(ctx)

	if errors.Is(err, db.ErrNotFound) {
		log.Logger.WithField("ids", ids).WithContext(ctx).Info("couldn't find team member moves by ids")
		return nil, fmt.Errorf("couldn't find team member moves by ids: %s", ids)
	}

	if err != nil {
		log.Logger.WithField("ids", ids).WithError(err).WithContext(ctx).Error("error getting team member moves by ids for deletion")
		return nil, fmt.Errorf("error getting team member moves by ids %s", ids)
	}

	_, err2 := r.client.TeamMemberMove.FindMany(db.TeamMemberMove.ID.In(ids)).Delete().Exec(ctx)

	if err2 != nil {
		log.Logger.WithField("ids", ids).WithError(err2).WithContext(ctx).Error("error deleting team member moves by ids")
		return nil, fmt.Errorf("error deleting team member moves by ids %s", ids)
	}

	moves := make([]*model.Move, 0)
	for _, result := range results {
		move := model.NewMoveFromDb(*result.PokemonMove().Move())
		moves = append(moves, &move)
	}
	return moves, nil
}
