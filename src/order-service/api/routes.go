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
	//not really a fan of this url pattern, add item with no (item id) might be better as im passing the Orderitem anyway
	//probably a better way to do this
	r.HandleFunc("/order/{orderId}/addItem/{itemId}", AddItemToOrder).Methods("POST")
	return r
}
