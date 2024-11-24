package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName   string             `bson:"first_name" json:"first_name" binding:"required"`
	LastName    string             `bson:"last_name" json:"last_name" binding:"required"`
	Email       string             `bson:"email" json:"email" binding:"required,email"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number" binding:"required"`
	Position    string             `bson:"position" json:"position"`
	Department  string             `bson:"department" json:"department"`
	DateOfHire  string             `bson:"date_of_hire" json:"date_of_hire"`
}
