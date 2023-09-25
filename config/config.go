package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importe o driver PostgreSQL
	"log"
)

var (
	db *sql.DB
)

// OpenDB efetua a conexão com o banco de dados
func OpenDB() {

	connStr := "user=postgres dbname=postgres password=1997 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")
}

// Retorna o ponteiro do banco de dadoss
func GetDB() *sql.DB {
	return db
}
