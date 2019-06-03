package router

import (
	"countryes/controllers/phoneCodeController"
	"countryes/controllers/reloadController"
	"countryes/middleware"

	"github.com/gorilla/mux"
)

func Startup() (*mux.Router, error) {
	router := mux.NewRouter()
	md := middleware.New()

	router.HandleFunc("/reload", md.PanicWrap(md.Log(reloadController.ReloadController))).Methods("POST")
	router.HandleFunc("/code/{name}", md.PanicWrap(md.Log(phoneCodeController.PhoneCodeController))).Methods("GET")

	return router, nil
}
