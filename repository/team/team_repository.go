package team

import "database/sql"

type teamRepository struct {
	database *sql.DB
}
