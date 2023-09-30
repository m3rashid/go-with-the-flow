package search

import "go.mongodb.org/mongo-driver/bson/primitive"

const RESOURCE_MODEL_NAME = "resources"

type Resource struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Name         string             `json:"name" bson:"name" validate:"required"`
	Description  string             `json:"desc" bson:"desc" validate:""`
	ResourceID   primitive.ObjectID `json:"rId" bson:"rId" validate:"required"`
	ResourceType string             `json:"rType" bson:"rType" validate:"required"`
}

// Not a model but a helper struct for the search module
type ResourceIndex struct {
	NameKey        string
	DescriptionKey string
	DisplayUrl     string
}
