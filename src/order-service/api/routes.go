package api

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/item", CreateItem).Methods("POST")
	//personally would separate the below route to /items, but done as test requests
	r.HandleFunc("/item", ListItems).Methods("GET")
	r.HandleFunc("/item/{id}", ListItems).Methods("GET")
	r.HandleFunc("/order", CreateOrder).Methods("POST")
	r.HandleFunc("/order/{id}/addItem{id}", AddItemToOrder).Methods("POST")
	return r
}
