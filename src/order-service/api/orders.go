package api

type Item struct {
	//would use uuid.UUID here for database, but also a "Create" function initializing uuid.v4 for record
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Items struct {
	Items []Item `json:"items"`
}

type Order struct {
	//would use uuid.UUID here for database, but also a "Create" function initializing uuid.v4 for record
	Id         string      `json:"id"`
	OrderItems []OrderItem `json:"order_items"`
	// TODO: add Get price of item function to calculate below
	Total  float64 `json:"total"`
	Status string  `json:"status"`
}

type OrderItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

//to replicate DB
var items = []Item{
	{
		Id:    "1",
		Name:  "burger",
		Price: 1.99,
	},
	{
		Id:    "2",
		Name:  "chips",
		Price: 0.99,
	},
}

//to replicate DB
var orders []Order

//to put in consts
var itemList = []string{"burger", "chips"}
