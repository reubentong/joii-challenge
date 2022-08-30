package api

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/item", CreateItem).Methods("POST")
	r.HandleFunc("/item", ListItem).Methods("GET")
	return r
}
