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

func GetSpecificOrderInfo(customerID, orderID string) string {
	models.CustomerList.Syncer.Lock()
	if models.CustomerList.List[customerID] == nil {
		models.CustomerList.Syncer.Unlock()
		return ""
	}
	models.CustomerList.Syncer.Unlock()
	return models.CustomerList.GetStatus(customerID, orderID).String()
}

func GetAllOrderStatus(customerID string) []models.CustomerAllOrderStatus {
	models.CustomerList.Syncer.Lock()
	defer models.CustomerList.Syncer.Unlock()
	if models.CustomerList.List[customerID] == nil {
		return nil
	}
	return convertToReponse(models.CustomerList.List[customerID])
}

func convertToReponse(orderStatus models.CustomerOrderStatus) []models.CustomerAllOrderStatus {
	allOrderStatus := make([]models.CustomerAllOrderStatus, 0)
	for key, value := range orderStatus {
		allOrderStatus = append(allOrderStatus, models.CustomerAllOrderStatus{
			OrderID:     key,
			OrderStatus: value.String(),
		})
	}
	return allOrderStatus
}
