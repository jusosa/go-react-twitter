package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jusosa/go-react-twitter/models"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	key := []byte("learning_go")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"last_name": user.LastName,
		"birthdate": user.BirthDate,
		"avatars":    user.Avatar,
		"banners":    user.Banner,
		"biography": user.BirthDate,
		"location":  user.Location,
		"web":       user.Web,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
