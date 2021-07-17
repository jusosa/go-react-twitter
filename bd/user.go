package bd

import (
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user models.User) (string, bool, error) {
	user.Password, _ = EncryptPassword(user.Password)
	return Create(user, "users")
}

func UserExists(email string) (models.User, bool, string) {
	condition := bson.M{"email": email}
	var user models.User

	err := FindOne(condition, "users").Decode(&user)
	ID := user.ID.Hex()

	status := true
	if err != nil {
		status = false
	}
	return user, status, ID
}
