package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct { //defining user type struct
	Id     bson.ObjectId `json:"id" bson:"_id"` //telling mongo db that in Json it will be id but in bson it will be _id  as ref to mongo db
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
