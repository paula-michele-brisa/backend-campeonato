package user

import (
	"github.com/paula-michele-brisa/backend-campeonato/config/logger"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	user3 "github.com/paula-michele-brisa/backend-campeonato/model/repository/convert/user"
	user2 "github.com/paula-michele-brisa/backend-campeonato/model/repository/entity/user"
	"github.com/paula-michele-brisa/backend-campeonato/model/user"
)

func (user *userRespository) FindUserByEmail(email string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	db := user.databaseConnection

	query := `SELECT email FROM t_user WHERE email=?`

	userEntity := &user2.UserEntity{}

	err := db.QueryRow(query, email).Scan(&userEntity)

	if userEntity == nil {
		logger.Info("Email não cadastrado")
	}

	if err != nil {

		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
	}

	return user3.ConvertEntityToDomain(*userEntity), nil
}

func (user *userRespository) FindUserByID(id string) (user.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository")

	db := user.databaseConnection

	query := `SELECT id FROM t_user WHERE id=?`

	userEntity := &user2.UserEntity{}

	err := db.QueryRow(query, id).Scan(&userEntity)

	if userEntity == nil {
		logger.Info("Usuário não cadastrado")
	}

	if err != nil {

		logger.Error("Ocorreu um erro ao tentar obter usuário no banco de dados", err)
	}

	return user3.ConvertEntityToDomain(*userEntity), nil
}
