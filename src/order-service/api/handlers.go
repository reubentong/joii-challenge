package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body, %v", err)
		w.WriteHeader(500)
		return
	}
	var item Item

	err = json.Unmarshal(requestBody, &item)
	if err != nil {
		log.Printf("Error unmashalling request, %v", err)
		w.WriteHeader(500)
		return
	}

	if contains(itemList, item.Name) {
		items = append(items, item)
		//not sure if this will be required for the post request, however to view for testing might be helpful
		jsonResponse, err := json.Marshal(items)
		if err != nil {
			log.Printf("Error marshalling response, %v", err)
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	} else {
		log.Printf("item name not allowed")
		w.WriteHeader(400)
		return
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
