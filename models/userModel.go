package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct{
	ID        primitive.ObjectID
	firstName string
	lastName string
	email string
	password string
	token string
	refresh_token string

}