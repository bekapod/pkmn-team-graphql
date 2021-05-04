package model

import "bekapod/pkmn-team-graphql/data/db"

type TeamMember struct {
	ID        string              `json:"id"`
	Slot      int                 `json:"slot"`
	PokemonID string              `json:"pokemonId"`
	Moves     *TeamMemberMoveList `json:"moves"`
	Team      *Team               `json:"team"`
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

	if dbTeam := dbTeamMember.RelationsTeamMember.Team; dbTeam != nil {
		team := NewTeamFromDb(*dbTeam)
		teamMember.Team = &team
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

func (l *TeamMemberList) RemoveTeamMember(id string) {
	teamMembers := make([]*TeamMember, 0)

	for _, member := range l.TeamMembers {
		if member.ID != id {
			teamMembers = append(teamMembers, member)
		}
	}

	l.Total = len(teamMembers)
	l.TeamMembers = teamMembers
}
