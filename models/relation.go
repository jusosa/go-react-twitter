package models

type Relation struct {
	UserId string `bson:"user_id" json:"user_id"`
	FollowingUser string `bson:"following_user" json:"following_user"`
}

type RelationResponse struct {
	Status bool `json:"status"`
}