package kitchen

import (
	"fmt"

	"github.com/Omkarz7/pizza-hut/models"
)

func MakePizza(tableOrder models.PizzaOrder) {
	models.CustomerList[tableOrder.CustomerID][tableOrder.OrderID] = models.PreparingDough
	prepareDough()
	models.CustomerList[tableOrder.CustomerID][tableOrder.OrderID] = models.AddingIngredients
	addIngredients()
	models.CustomerList[tableOrder.CustomerID][tableOrder.OrderID] = models.Baking
	bakePizza()

	models.CustomerList[tableOrder.CustomerID][tableOrder.OrderID] = models.Ready
	models.NotifyStatus <- fmt.Sprintf("Dear Customer %s, your Order %s is Ready\n", tableOrder.CustomerID, tableOrder.OrderID)
}
