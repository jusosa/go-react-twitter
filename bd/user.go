package bd

import (
	"context"
	"fmt"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func UpdateUser(user models.User, ID string) (bool, error) {
	newData := make(map[string]interface{})
	if len(user.Name) > 0 {
		newData["name"] = user.Name
	}
	if len(user.LastName) > 0 {
		newData["last_name"] = user.LastName
	}
	if len(user.Avatar) > 0 {
		newData["avatars"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		newData["banners"] = user.Banner
	}
	if len(user.Biography) > 0 {
		newData["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		newData["location"] = user.Location
	}
	if len(user.Web) > 0 {
		newData["web"] = user.Web
	}
	if !user.BirthDate.IsZero() {
		newData["birthdate"] = user.BirthDate
	}

	updateStr := bson.M{
		"$set": newData,
	}

	objId, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objId}}

	_, err := UpdateOne(filter, updateStr, "users")
	if err != nil {
		return false, err
	}

	return true, nil
}

func FindAllUsers(Id string, page int64, search string, typeUser string) ([] *models.User, bool) {
	var results [] *models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := FindAllByCondition(query, findOptions, "users")

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool
	transactionContext := context.TODO()

	for cursor.Next(transactionContext) {
		var user models.User
		err = cursor.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relation models.Relation
		relation.UserId = Id
		relation.FollowingUser = user.ID.Hex()

		include = false

		found, err = FindRelation(relation)
		if typeUser == "new" && !found {
			include = true
		}

		if typeUser == "follow" && found {
			include = true
		}

		if relation.FollowingUser == Id {
			include = false
		}

		if include {
			user.Password = ""
			user.Biography = ""
			user.Web = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(transactionContext)

	return results, true
}
