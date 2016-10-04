package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Object struct {
	ObjectID   bson.ObjectId `bson:"_id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
