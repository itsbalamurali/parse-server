package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Object
	Username string
	Password string
	Email string
	Extra     bson.M `bson:",inline"`
}
