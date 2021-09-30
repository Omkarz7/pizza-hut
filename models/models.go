package models

// PizzaOrder srtuct gets pizza order from the customer
type PizzaOrder struct {
	CustomerID string `json:"customerID"`
	PizzaSize  string `json:"pizzaSize"`
	PizzaType  string `json:"pizzaType"`
	OrderID    string
}

//Configs for host and port
var host = "http://localhost"
var Port = ":4700"

var NotifyStatus = make(chan string, 10)
