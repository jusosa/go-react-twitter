package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tweet struct {
	ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string `bson:"user_id" json:"user_id,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	CreationDate time.Time `bson:"creation_date" json:"creation_date,omitempty"`
}

type TweetBody struct {
	Message string `bson:"message" json:"message"`
}