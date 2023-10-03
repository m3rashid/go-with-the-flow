package flow

import "go.mongodb.org/mongo-driver/bson/primitive"

type Node struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id" validate:""`
	Name       string             `json:"name" bson:"name" validate:"required"`
	Collection string             `json:"collection" bson:"collection" validate:"required"`

	// Trigger can be insert, update, delete or replace
	Trigger string `json:"trigger" bson:"trigger" validate:"required,omitempty"`

	// Handler is a function that takes in the data document/id and returns it with possible modifications
	Handler string `json:"handler" bson:"handler" validate:"required,omitempty"`

	// a function that returns a node index, defaults to () => currentNode + 1
	NextNode string `json:"nextNode" bson:"nextNode" validate:"required,omitempty"`
}
