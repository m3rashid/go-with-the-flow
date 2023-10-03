package permission

import (
	"github.com/m3rashid/go-with-the-flow/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ROLE_MODEL_NAME = "roles"

type Role struct {
	db.BaseSchema `bson:",inline"`
	Name          string                 `json:"name" bson:"name" validate:"required"`
	Description   string                 `json:"description" bson:"description" validate:""`
	Permissions   [](primitive.ObjectID) `json:"permissions" bson:"permissions" validate:""`
}
