package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"last_name" json:"last_name,omitempty"`
	BirthDate time.Time          `bson:"birthdate" json:"birthdate,omitempty"`
	Email     string             `bson:"mail" json:"mail"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatars" json:"avatars,omitempty"`
	Banner    string             `bson:"banners" json:"banners,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	Web       string             `bson:"web" json:"web,omitempty"`
}
