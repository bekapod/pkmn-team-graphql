package model

type TeamMember struct {
	ID        string              `json:"id"`
	Slot      int                 `json:"slot"`
	PokemonID string              `json:"pokemonId"`
	Moves     *TeamMemberMoveList `json:"moves"`
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
