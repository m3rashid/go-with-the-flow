package auth

import "github.com/m3rashid/go-server/db"

const PROFILE_MODEL_NAME = "profiles"

type Profile struct {
	db.BaseSchema `bson:",inline"`
}
