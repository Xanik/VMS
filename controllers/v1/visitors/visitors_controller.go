package visitorscontoller

import (
	"net/http"
	"time"
	dbs "vms/config/db"
	"vms/config/responses"
	"vms/libs/files"
	httplib "vms/libs/http"

	"gopkg.in/mgo.v2/bson"

	visitorsmodel "vms/models/visitors"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

var (
	env          = viper.GetString("env")
	dbName       = "vms"
	dbCollection = "visitors"
)

//UploadImage controller
func UploadImage(res http.ResponseWriter, req *http.Request) {
	url := files.UploadFile("image", req)

	resp := responses.GeneralResponse{Success: true, Data: url, Message: "image uploaded"}

	httplib.Response(res, resp)
}

//RegisterVistor controller
func RegisterVistor(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	var data visitorsmodel.Visitors

	c.BindJSON(&data)

	data.ID = bson.NewObjectId()
	data.Date = time.Now()

	err := coll.Insert(data)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating visitor"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: data, Message: "visitor created"}
	httplib.Response(res, resp)
}

//GetVisitorDetailsByEmail controller
func GetVisitorDetailsByEmail(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	visitorEmail := c.Params("visitorEmail")

	var visitor interface{}

	err := coll.Find(bson.M{"email": visitorEmail}).One(&visitor)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error getting visitor"}
		httplib.Response(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: visitor, Message: "visitor details"}
	httplib.Response(res, resp)
}

//DeleteVistor controller
func DeleteVistor(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	visitorEmail := c.Params("visitorEmail")

	err := coll.Remove(bson.M{"email": visitorEmail})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error deleting visitor"}
		httplib.Response(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: visitorEmail, Message: "visitor deleted"}
	httplib.Response(res, resp)
}

//UpdateVistorDetials controller
func UpdateVistorDetials(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}
	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	var updates bson.M

	c.BindJSON(updates)
	visitorID := c.Params("visitorEmail")

	err := coll.Update(bson.M{"_id": bson.ObjectIdHex(visitorID)}, bson.M{"$set": updates})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error updating visitor"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: updates, Message: "visitor updated"}
	httplib.Response(res, resp)
}
