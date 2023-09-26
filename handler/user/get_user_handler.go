package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paula-michele-brisa/backend-campeonato/config"
	"log"
)

// GetUserHandler é responsável por retorar os dados do usuário
func GetUserHandler(context *gin.Context) {
	database := config.GetDB()

	rows, err := database.Query("SELECT * FROM pg_user")

	defer rows.Close()

	for rows.Next() {
		var name string
		var email string
		var password string

		if err := rows.Scan(&name, &email, &password); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Name: %s, Email: %s, Password: %s", name, email, password)
	}

	if err != nil {
		log.Fatal(err)

	}

	context.JSON(200, gin.H{
		"usario": "aquis",
	})
}
