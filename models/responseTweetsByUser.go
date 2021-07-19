package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ResponseTweetsByUser struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
	UserId string `bson:"user_id" json:"user_id"`
	FollowingUser string `bson:"following_user" json:"following_user"`
	Tweet struct{
		Message string `bson:"message" json:"message"`
		Date time.Time `bson:"date" json:"date"`
		ID string `bson:"_id" json:"_id"`
	} `bson:"tweet" json:"tweet"`
}