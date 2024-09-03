package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username" validate:"required,min=3,max=32"`
	Email    string             `bson:"email" validate:"required,email"`
	Password string             `bson:"password" validate:"required,min=6"`
}
