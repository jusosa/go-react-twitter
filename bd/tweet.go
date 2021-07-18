package bd

import (
	"context"
	"github.com/jusosa/go-react-twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateTweet(tweet models.Tweet) (string, bool, error) {
	return Create(tweet, "tweet")
}

func FindTweets(ID string, page int64) ([]*models.Tweet, bool) {
	var result []*models.Tweet
	condition := bson.M{
		"user_id": ID,
	}

	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "creation_date", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	cursor, err := FindAllByCondition(condition, opts, "tweet")
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var t models.Tweet
		err := cursor.Decode(&t)
		if err != nil {
			return result, false
		}
		result = append(result, &t)
	}

	return result, true
}

func DeleteTweet(Id string, UserId string) error {
	objId, _ := primitive.ObjectIDFromHex(Id)
	condition := bson.M{
		"_id":     objId,
		"user_id": UserId,
	}

	err := DeleteOne(condition, "tweet")
	return err
}
