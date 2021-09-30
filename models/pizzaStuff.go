package models

// PreparationTime stores time required to prepare certain steps in seconds
type PreparationTime struct {
	PrepareDough   int
	AddIngredients int
	BakePizza      int
}

type OrderStatus int

const (
	Recieved OrderStatus = iota
	PreparingDough
	AddingIngredients
	Baking
	Ready
)

func (status OrderStatus) String() string {
	return [...]string{"Recieved", "PreparingDough", "AddingIngredients", "Baking", "Ready"}[status]
}

var PizzaPreparationTime = PreparationTime{PrepareDough: 5, AddIngredients: 5, BakePizza: 15}

type CustomerOrderStatus map[string]OrderStatus

var CustomerList = make(map[string]CustomerOrderStatus)
