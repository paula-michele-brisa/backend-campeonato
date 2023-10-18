package login

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paula-michele-brisa/backend-campeonato/config/rest_err"
	"os"
	"time"
)

var (
	JWT_SECRET_KEY = "JWT_SECRETY_KEY"
)

func (login *loginDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"ID":    login.ID,
		"email": login.Email,
		"name":  login.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(err.Error())
	}

	return tokenString, nil
}
