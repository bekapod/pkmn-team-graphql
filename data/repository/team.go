package repository

import "bekapod/pkmn-team-graphql/data/db"

type Team struct {
	client *db.PrismaClient
}

func NewTeam(client *db.PrismaClient) Team {
	return Team{
		client: client,
	}
}
