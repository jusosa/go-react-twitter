package bd

import (
	"fmt"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (string, bool, error) {
	user.Password, _ = EncryptPassword(user.Password)
	return Create(user, "users")
}

func UserExists(email string) (models.User, bool, string) {
	condition := bson.M{"mail": email}
	var user models.User

	err := FindOne(condition, "users").Decode(&user)
	ID := user.ID.Hex()

	status := true
	if err != nil {
		status = false
	}
	return user, status, ID
}

func TryLogin(mail string, password string) (models.User, bool) {
	user, exists, _ := UserExists(mail)

	if !exists {
		return user, exists
	}

	passwordEncrypted := []byte(password)
	passwordEncryptedDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordEncryptedDB, passwordEncrypted)
	if err != nil {
		return user, false
	}

	return user, true
}

func ViewProfile(ID string) (models.User, error) {
	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{"_id": objID}
	var profile models.User

	err := FindOne(condition, "users").Decode(&profile)
	profile.Password = ""

	if err != nil {
		fmt.Println("Profile not found: " + err.Error())
		return profile, err
	}

	return profile, nil
}
