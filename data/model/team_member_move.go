package model

type TeamMemberMove struct {
	ID            string `json:"id"`
	Slot          int    `json:"slot"`
	PokemonMoveID string `json:"pokemonMoveId"`
}

func NewTeamMemberMoveList(teams []*TeamMemberMove) TeamMemberMoveList {
	return TeamMemberMoveList{
		Total:           len(teams),
		TeamMemberMoves: teams,
	}
}

func NewEmptyTeamMemberMoveList() TeamMemberMoveList {
	return TeamMemberMoveList{
		Total:           0,
		TeamMemberMoves: []*TeamMemberMove{},
	}
}

func (l *TeamMemberMoveList) AddTeamMemberMove(tm *TeamMemberMove) {
	l.Total++
	l.TeamMemberMoves = append(l.TeamMemberMoves, tm)
}
