package model

import "bekapod/pkmn-team-graphql/data/db"

type TeamMember struct {
	ID        string              `json:"id"`
	Slot      int                 `json:"slot"`
	PokemonID string              `json:"pokemonId"`
	Moves     *TeamMemberMoveList `json:"moves"`
}

func NewTeamMemberFromDb(dbTeamMember db.TeamMemberModel) TeamMember {
	teamMemberMoves := NewEmptyTeamMemberMoveList()
	teamMember := TeamMember{
		ID:        dbTeamMember.ID,
		Slot:      dbTeamMember.Slot,
		PokemonID: dbTeamMember.PokemonID,
		Moves:     &teamMemberMoves,
	}

	for _, tmm := range dbTeamMember.Moves() {
		teamMemberMove := NewTeamMemberMoveFromDb(tmm)
		teamMember.Moves.AddTeamMemberMove(&teamMemberMove)
	}

	return teamMember
}

func NewTeamMemberList(teams []*TeamMember) TeamMemberList {
	return TeamMemberList{
		Total:       len(teams),
		TeamMembers: teams,
	}
}

func NewEmptyTeamMemberList() TeamMemberList {
	return TeamMemberList{
		Total:       0,
		TeamMembers: []*TeamMember{},
	}
}

func (l *TeamMemberList) AddTeamMember(tm *TeamMember) {
	l.Total++
	l.TeamMembers = append(l.TeamMembers, tm)
}
