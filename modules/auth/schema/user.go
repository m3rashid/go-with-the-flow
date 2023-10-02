package auth

import (
	"github.com/m3rashid/server/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USER_MODEL_NAME = "users"

type User struct {
	db.BaseSchema `bson:",inline"`
	Name          string                 `json:"name" bson:"name" validate:"required"`
	Email         string                 `json:"email" bson:"email" validate:"required,email"`
	Phone         string                 `json:"phone" bson:"phone" validate:""`
	Avatar        string                 `json:"avatar" bson:"avatar" validate:""`
	Deactivated   bool                   `json:"deactivated" bson:"deactivated" validate:""`
	Password      string                 `json:"password" bson:"password" validate:"required"`
	Roles         [](primitive.ObjectID) `json:"roles" bson:"role,omitempty" validate:""`
	Profile       primitive.ObjectID     `json:"profile" bson:"profile,omitempty" validate:""`
}
