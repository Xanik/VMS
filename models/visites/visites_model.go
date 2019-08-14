package visitesmodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Visties model
type Visties struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	Visitor  bson.ObjectId `json:"visitor_id" bson:"visitor_id"`
	Type     string        `json:"type" bson:"type"`
	Purpose  string        `json:"purpose" bson:"purpose"`
	Employee bson.ObjectId `json:"employee_id" bson:"employee_id"`
	Code     string        `json:"code" bson:"code"`
	Status   string        `json:"status" bson:"status"`
	Date     time.Time     `json:"date" bson:"date"`
}
