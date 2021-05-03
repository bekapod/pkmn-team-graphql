package model

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
