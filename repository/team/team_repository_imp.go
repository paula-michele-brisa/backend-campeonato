package team

import "database/sql"

func TeamRepository(database *sql.DB) TeamRepositoryInterface {
	return &teamRepository{
		databaseConnection: database,
	}
}
