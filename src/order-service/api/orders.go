package api

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

//to replicate DB
var items []Item

var itemList = []string{"burger", "chips"}
