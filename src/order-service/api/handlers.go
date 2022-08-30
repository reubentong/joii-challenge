package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type ErrorHandler interface {
	APIError() (int, string)
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}

func JSONHandleError(w http.ResponseWriter, err error) {
	var apiErr ErrorHandler
	if errors.As(err, &apiErr) {
		status, msg := apiErr.APIError()
		JSONError(w, msg, status)
	} else {
		JSONError(w, "internal server error", http.StatusInternalServerError)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	//i wouldnt usually add all these errors, however for completeness...
	if err != nil {
		log.Printf("error reading body, %v", err)
		JSONHandleError(w, err)
		return
	}
	var item Item

	err = json.Unmarshal(requestBody, &item)
	if err != nil {
		log.Printf("error unmashalling request, %v", err)
		JSONHandleError(w, err)
		return
	}

	if containsInDb(items, item.Id) {
		log.Printf("id already exists")
		JSONHandleError(w, err)
		return
	}

	if contains(itemList, item.Name) {
		items = append(items, item)
		//not sure if a response will be required for the post request, however to view for testing might be helpful
		jsonResponse, err := json.Marshal(items)
		if err != nil {
			log.Printf("error marshalling response, %v", err)
			JSONHandleError(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jsonResponse)
		if err != nil {
			log.Printf("error writing response response, %v", err)
			JSONHandleError(w, err)
			return
		}
	} else {
		log.Printf("item name not allowed")
		JSONHandleError(w, err)
		return
	}
}

func ListItems(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var response Items

	if id != "" {
		for _, item := range items {
			if item.Id == id {
				response.Items = append(response.Items, item)
			}
		}
	} else {
		response.Items = items
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("error marshalling response, %v", err)
		JSONHandleError(w, err)
		return
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("error writing response response, %v", err)
		JSONHandleError(w, err)
		return
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body, %v", err)
		JSONHandleError(w, err)
		return
	}
	var order Order

	err = json.Unmarshal(requestBody, &order)
	if err != nil {
		log.Printf("error unmashalling request, %v", err)
		JSONHandleError(w, err)
		return
	}

	var total float64
	for _, item := range order.OrderItems {
		total = total + GetPrice(item.Name)*float64(item.Quantity)
	}
	order.Total = total

	orders = append(orders, order)
	//not sure if a response will be required for the post request, however to view for testing might be helpful
	jsonResponse, err := json.Marshal(orders)
	if err != nil {
		log.Printf("error marshalling response, %v", err)
		JSONHandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("error writing response response, %v", err)
		JSONHandleError(w, err)
		return
	}
}

func AddItemToOrder(w http.ResponseWriter, r *http.Request) {

}

// GetPrice to grab the price of item
func GetPrice(name string) (price float64) {
	for _, item := range items {
		if item.Name == name {
			price = item.Price
		}
	}
	return
}

//from google https://freshman.tech/snippets/go/check-if-slice-contains-element/
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func containsInDb(s []Item, str string) bool {
	for _, v := range s {
		if v.Id == str {
			return true
		}
	}
	return false
}
