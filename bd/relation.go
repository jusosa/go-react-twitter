package bd

import (
	"context"
	"fmt"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateRelation(relation models.Relation) (string, bool, error) {
	return Create(relation, "relation")
}

func DeleteRelation(relation models.Relation) error {
	return DeleteOne(relation, "relation")
}

func FindRelation(relation models.Relation) (bool, error) {
	condition := bson.M{
		"user_id":        relation.UserId,
		"following_user": relation.FollowingUser,
	}

	var result models.Relation

	err := FindOne(condition, "relation").Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}

func FindTweetsByUser(ID string, page int) ([]models.ResponseTweetsByUser, bool) {
	skip := (page - 1) * 20

	query := make([]bson.M, 0)
	query = append(query, bson.M{"$match": bson.M{"user_id": ID}})
	query = append(query, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "following_user",
			"foreignField": "user_id",
			"as":           "tweet",
		},
	})
	query = append(query, bson.M{"$unwind": "$tweet"})
	query = append(query, bson.M{"$sort": bson.M{"tweet.date": -1}})
	query = append(query, bson.M{"$skip": skip})
	query = append(query, bson.M{"$limit": 20})

	cursor, err := Aggregate(query, "relation")
	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}

	var results []models.ResponseTweetsByUser
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	return results, true
}
