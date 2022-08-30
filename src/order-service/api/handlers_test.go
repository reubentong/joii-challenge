package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateItem(t *testing.T) {
	newItem := Item{
		Id:    "3",
		Name:  "burger",
		Price: 1.99,
	}

	body, err := json.Marshal(newItem)
	if err != nil {
		t.Errorf("cannot marshall request body")
	}

	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("error making new request w/body: %d", req.Body)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(CreateItem).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code error. wanted %d, got %d", http.StatusOK, status)
	}

	data, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	var returnedBody []Item
	if err := json.Unmarshal(data, &returnedBody); err != nil {
		t.Errorf("returned list of items is invalid JSON. Got: %s", data)
	}
}

func TestListItems(t *testing.T) {
	req, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(ListItems).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code error. wanted %d, got %d", http.StatusOK, status)
	}

	data, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		t.Fatal(err)
	}

	var responseItems Items

	if err := json.Unmarshal(data, &responseItems); err != nil {
		t.Errorf("returned list is invalid JSON. Got: %s", data)
	}
	for index, item := range responseItems.Items {
		if item != items[index] {
			t.Errorf("returned list is not the same as real. Expected %+v. Got %+v instead", items[index], responseItems.Items)
		}
	}
}
