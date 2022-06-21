package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Writer      User               `json:"writer,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	CreatedAt   primitive.DateTime `json:"createdAt,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updatedAt,omitempty"`
}
