package config

import (
	"database/sql"
	_ "github.com/lib/pq" // Importe o driver PostgreSQL
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"os"
)

var (
	PostgresCredential = "URL_POSTGRES"
)

// NewDatabasePostgres efetua a conexão com o banco de dados
func NewDatabasePostgres() (*sql.DB, error) {

	url_postgress := os.Getenv(PostgresCredential)

	db, err := sql.Open("postgres", url_postgress)

	if err != nil {
		logger.Error("Erro ao abrir a conexão com o PostgreSQL:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		logger.Error("Erro na conexão com o PostgreSQL:", err)
	}

	return db, nil
}
