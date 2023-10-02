package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActionLog struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Time time.Time          `json:"time,omitempty" bson:"time,omitempty"`
}

type BaseSchema struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:""`
	Deleted   bool               `json:"deleted,omitempty" bson:"deleted,omitempty" validate:""`
	CreatedBy *ActionLog         `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedBy *([]ActionLog)     `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
}
