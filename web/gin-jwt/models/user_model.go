package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    Id       primitive.ObjectID     `json:"id,omitempty"`
    username    string             `json:"username,omitempty" validate:"required"`
    password    string             `json:"password,omitempty" validate:"required"`
}