package models


import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Bill struct{
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	Created_by *string             `json:"created_by" validate:"required,min=3,max=100"`
	Customername *string              `json:"customer_name" validate:"required,min=3,max=100"`
	Items map[string]float64       `json:"items" validate:"required"`
	Email *string                  `json:"email" validate:"email"`
	Created_at   time.Time         `json:"created_at"`
	Updated_at   time.Time         `json:"updated_at"`
	Bill_ID string                 `json:"bill_id"`


}