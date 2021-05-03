package model

import "bekapod/pkmn-team-graphql/data/db"

func NewTeamFromDb(dbTeam db.TeamModel) Team {
	teamMembers := NewEmptyTeamMemberList()
	team := Team{
		ID:      dbTeam.ID,
		Name:    dbTeam.Name,
		Members: &teamMembers,
	}

	for _, tm := range dbTeam.TeamMembers() {
		teamMember := NewTeamMemberFromDb(tm)
		teamMembers.AddTeamMember(&teamMember)
	}

	return team
}

func NewTeamList(teams []*Team) TeamList {
	return TeamList{
		Total: len(teams),
		Teams: teams,
	}
}

func NewEmptyTeamList() TeamList {
	return TeamList{
		Total: 0,
		Teams: []*Team{},
	}
}

func (l *TeamList) AddTeam(t *Team) {
	l.Total++
	l.Teams = append(l.Teams, t)
}
