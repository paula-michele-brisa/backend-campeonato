package teams

import (
	"github.com/gin-gonic/gin"
)

// GetTeamHandler obtém um time pelo id
func (team *teamHandler) FindTeamByIDHandler(context *gin.Context) {
	//var res string
	//var teams []string
	//rows, err := db.Query("SELECT * FROM teams ")
	//defer rows.Close()
	//
	//if err != nil {
	//	log.Fatal(err)
	//	context.JSON(500, "Ocorreu um erro ao tentar inserir daddo na tabela Team")
	//
	//	for rows.Next() {
	//		rows.Scan(&res)
	//		teams = append(teams, res)
	//	}
	//}

	context.JSON(200, gin.H{
		"total": "10",
	})
}
