package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jusosa/go-react-twitter/bd"
	"github.com/jusosa/go-react-twitter/models"
	"strings"
)

var Email string
var UserId string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("learning_go")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("token format is invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, founded, _ := bd.UserExists(claims.Email)
		if founded {
			Email = claims.Email
			UserId = claims.ID.Hex()
		}
		return claims, founded, UserId, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("invalid token")
	}
	return claims, false, "", err
}
