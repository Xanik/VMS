package visitescontroller

import (
	"net/http"
	"time"
	dbs "vms/config/db"
	"vms/config/responses"
	httplib "vms/libs/http"
	visitesmodel "vms/models/visites"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	env          = viper.GetString("env")
	dbName       = "vms"
	dbCollection = "visitors"
)

func generateCheckoutCode(id string) string {
	code := id[20:]
	return code
}

//CreateVisits controller
func CreateVisits(res http.ResponseWriter, req *http.Request) {
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

	var data visitesmodel.Visties

	c.BindJSON(&data)

	data.ID = bson.NewObjectId()
	data.Date = time.Now()
	data.Status = "checkin"
	data.Code = generateCheckoutCode(data.ID.Hex())

	err := coll.Insert(data)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating visit"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: data, Message: "visit created"}
	httplib.Response(res, resp)
}

//GetVisitorVisites controller
func GetVisitorVisites(res http.ResponseWriter, req *http.Request) {
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

	var visits []interface{}

	err := coll.Find(bson.M{"email": visitorEmail}).All(&visits)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error getting visitor visits"}
		httplib.Response(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: visits, Message: "visitor visits"}
	httplib.Response(res, resp)
}

//GetEmployeeVisites controller
func GetEmployeeVisites(res http.ResponseWriter, req *http.Request) {
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

	employeeEmail := c.Params("employeeEmail")

	var visits []interface{}

	err := coll.Find(bson.M{"email": employeeEmail}).All(&visits)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error getting employee visits"}
		httplib.Response(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: visits, Message: "employee visits"}
	httplib.Response(res, resp)
}

//Checkout controller
func Checkout(res http.ResponseWriter, req *http.Request) {
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
	visitorCode := c.Params("visitorCode")

	err := coll.Update(bson.M{"code": visitorCode}, bson.M{"$set": bson.M{"status": "checkout"}})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error updating visitor"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: updates, Message: "visitor updated"}
	httplib.Response(res, resp)
}
