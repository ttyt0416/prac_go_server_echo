package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Password  string             `json:"password,omitempty" validate:"required"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty"`
}
