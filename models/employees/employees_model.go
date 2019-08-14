package employeesmodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Employees model
type Employees struct {
	ID        bson.ObjectId `json:"_id" bson:"_id"`
	FirstName string        `json:"first_name" bson:"first_name"`
	LastName  string        `json:"last_name" bson:"last_name"`
	Email     string        `json:"email" bson:"emaol"`
	Image     string        `json:"image" bson:"image"`
	Date      time.Time     `json:"date" bson:"date"`
}
