package routes

import (
	"net/http"
	"vms/config/responses"
	visitescontroller "vms/controllers/v1/visites"
	visitorscontoller "vms/controllers/v1/visitors"
	httplib "vms/libs/http"
	mws "vms/middlewares"

	"github.com/gorilla/mux"
)

//Router for all routes
func Router() *mux.Router {
	route := mux.NewRouter()

	//BASE ROUTE
	route.HandleFunc("/v1", func(res http.ResponseWriter, req *http.Request) {
		resp := responses.GeneralResponse{Success: true, Message: "vms  server running....", Data: "vsm SERVER v1.0"}
		httplib.Response(res, resp)
	})

	route.Use(mws.AccessLogToConsole)

	//************************
	// VISITES  ROUTES
	//************************
	visitsRoute := route.PathPrefix("/v1/visits").Subrouter()
	visitsRoute.HandleFunc("", visitescontroller.CreateVisits).Methods("POST")
	visitsRoute.HandleFunc("/{employeeEmail}", visitescontroller.GetEmployeeVisites).Methods("GET")
	visitsRoute.HandleFunc("/{visitorEmail}", visitescontroller.GetVisitorVisites).Methods("GET")
	visitsRoute.HandleFunc("/{visitorCode}", visitescontroller.Checkout).Methods("PUT")

	//************************
	// VISITORS  ROUTES
	//************************
	visitorsRoute := route.PathPrefix("/v1/visitors").Subrouter()
	visitorsRoute.HandleFunc("", visitorscontoller.RegisterVistor).Methods("POST")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.GetVisitorDetailsByEmail).Methods("GET")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.UpdateVistorDetials).Methods("PUT")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.UploadImage).Methods("POST")
	visitorsRoute.HandleFunc("/{visitorEmail}", visitorscontoller.DeleteVistor).Methods("DELETE")
	return route
}
