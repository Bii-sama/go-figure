package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct{
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Firstname *string              `json:"first_name" validate:"required,min=3,max=100"`
	Lastname *string               `json:"last_name"`
	Email *string                  `json:"email"`
	Password *string               `json:"password"`
	Token *string                  `json:"token"`
	Refresh_token *string          `json:"refresh_token"`
	Created_at   time.Time
	Updated_at   time.Time
	User_ID string

}