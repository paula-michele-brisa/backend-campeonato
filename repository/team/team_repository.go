package team

import "database/sql"

type teamRepository struct {
	database *sql.DB
}

func TeamRespository(database *sql.DB) TeamRepositoryInterface {
	return &teamRepository{
		database: database,
	}
}
