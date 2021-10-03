package models

import "sync"

// PreparationTime stores time required to prepare certain steps in seconds
type PreparationTime struct {
	PrepareDough   int
	AddIngredients int
	BakePizza      int
}

type OrderStatus int

const (
	NotFound OrderStatus = iota
	Recieved
	PreparingDough
	AddingIngredients
	Baking
	Ready
)

func (status OrderStatus) String() string {
	return [...]string{"Not found", "Recieved", "PreparingDough", "AddingIngredients", "Baking", "Ready"}[status]
}

var PizzaPreparationTime = PreparationTime{PrepareDough: 5, AddIngredients: 5, BakePizza: 15}

type CustomerOrderStatus map[string]OrderStatus

// var CustomerList = make(map[string]CustomerOrderStatus)

type CustomerOrderList struct {
	List   map[string]CustomerOrderStatus
	Syncer sync.Mutex
}

var CustomerList = CustomerOrderList{List: make(map[string]CustomerOrderStatus)}

func (c *CustomerOrderList) SetStatus(customerID, orderID string, status OrderStatus) {
	c.Syncer.Lock()
	defer c.Syncer.Unlock()
	c.List[customerID][orderID] = status
}

func (c *CustomerOrderList) GetStatus(customerID, orderID string) OrderStatus {
	c.Syncer.Lock()
	defer c.Syncer.Unlock()
	return c.List[customerID][orderID]
}
