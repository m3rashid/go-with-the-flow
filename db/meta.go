package db

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func OnCreate(doc *map[string]interface{}, doneBy primitive.ObjectID) error {
	if doneBy.String() == "" {
		return errors.New("no user found")
	}

	(*doc)["createdBy"] = struct {
		CreatedBy *ActionLog `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	}{
		&ActionLog{
			ID:   doneBy,
			Time: time.Now(),
		},
	}

	return nil
}

func OnUpdate(doc *map[string]interface{}, doneBy primitive.ObjectID) (interface{}, error) {
	if doneBy.String() == "" {
		return nil, errors.New("no user found")
	}

	change := bson.M{"$push": bson.M{"updatedBy": bson.M{
		"updatedBy": doneBy,
		"time":      time.Now(),
	}}}
	return change, nil
}
