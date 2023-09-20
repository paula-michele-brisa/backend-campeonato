package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Importe o driver PostgreSQL
	"log"
)

// OpenDB efetua a conexão com o banco de dados
func OpenDB() {
	db, err := sql.Open("postgres", "momonosuke")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")
}
