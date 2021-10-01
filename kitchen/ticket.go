package kitchen

import (
	"github.com/Omkarz7/pizza-hut/models"
)

func CustomerOrder(customerOrder models.PizzaOrder) {

	models.CustomerList.Syncer.Lock()
	if models.CustomerList.List[customerOrder.CustomerID] == nil {
		models.CustomerList.List[customerOrder.CustomerID] = make(models.CustomerOrderStatus)
	}
	models.CustomerList.Syncer.Unlock()

	models.CustomerList.SetStatus(customerOrder.CustomerID, customerOrder.OrderID, models.Recieved)
	MakePizza(customerOrder)
}
