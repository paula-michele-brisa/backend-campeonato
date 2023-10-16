package login

import (
	"database/sql"
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"github.com/paula-michele-brisa/backend-campeonato/domain/login"
	login2 "github.com/paula-michele-brisa/backend-campeonato/entity/login"
	user4 "github.com/paula-michele-brisa/backend-campeonato/repository/convert/user"
)

func (login *loginRepository) LoginRepository(email, password string) (login.LoginDomainInterface, *rest_err.RestErr) {
	db := login.databaseConnection

	query := `SELECT id, name, email FROM t_user WHERE email=$1 AND password=$2 `

	loginEntity := &login2.LoginEntity{}

	err := db.QueryRow(query, email, password).Scan(&loginEntity.ID, &loginEntity.Name, &loginEntity.Email)

	if err != nil {
		logger.Error("Ocorreu um erro ao tentar obter login no banco de dados", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return user4.ConvertLoginEntityToDomain(*loginEntity), nil
}

func LoginRepository(database *sql.DB) LoginRepositoryInterface {
	return &loginRepository{
		databaseConnection: database,
	}
}
