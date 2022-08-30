package main

import (
	"github.com/reubentong/order-service/api"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", api.Router()))
}
