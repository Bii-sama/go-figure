package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct{
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Firstname *string              `json:"first_name" validate:"required,min=3,max=100"`
	Lastname *string               `json:"last_name" validate:"required,min=3,max=100"`
	Email *string                  `json:"email" validate:"email, required"`
	Password *string               `json:"password" validate:"required,min=6,max=50"`
	Token *string                  `json:"token"`
	Refresh_token *string          `json:"refresh_token"`
	Created_at   time.Time         `json:"created_at"`
	Updated_at   time.Time         `json:"updated_at"`
	User_ID string                 `json:"user_id"`

}