package httplib

import (
	"encoding/json"
	"net/http"
	"vms/config/responses"

	"github.com/gorilla/mux"
)

//C serves as a place holder for the http request and response
type C struct {
	Req *http.Request
	Res http.ResponseWriter
}

//H defines a json type formate
type H map[string]interface{}

//BindJSON decodes http request body to a given object
func (c *C) BindJSON(data interface{}) {
	json.NewDecoder(c.Req.Body).Decode(data)
}

//JSON returns a http response encoded in application/json format to the response writer
func responseJSON(res http.ResponseWriter, status int, object interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(object)
}

//Params maps routes params to mux and returns the value of the key
func (c *C) Params(key string) string {
	return mux.Vars(c.Req)[key]
}

//Response returns a json response to the requester
func Response(res http.ResponseWriter, resp responses.GeneralResponse) {
	if resp.Success == false {
		responseJSON(res, 200, H{"error": resp.Error, "status": false, "message": resp.Message})
	} else {
		responseJSON(res, 200, H{"data": resp.Data, "status": true, "message": resp.Message})
	}
}

//Response400 returns a json response to the requester
func Response400(res http.ResponseWriter, resp responses.GeneralResponse) {
	if resp.Success == false {
		responseJSON(res, 400, H{"error": resp.Error, "status": false, "message": resp.Message})
	} else {
		responseJSON(res, 400, H{"data": resp.Data, "status": true, "message": resp.Message})
	}
}
