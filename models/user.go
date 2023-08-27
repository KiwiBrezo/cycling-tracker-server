package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserId   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Token    string             `json:"token"`
}
