package kitchen

import (
	"testing"
	"time"

	"github.com/Omkarz7/pizza-hut/models"
	"github.com/stretchr/testify/assert"
)

func TestPizzaOrdering(t *testing.T) {
	testOrder := models.PizzaOrder{
		CustomerID: "SomeID",
		PizzaType:  "VegPizza",
		PizzaSize:  "Medium",
		OrderID:    "SomeID",
	}

	go CustomerOrder(testOrder)

	time.Sleep(time.Second * 1)

	assert.Equal(t, models.PreparingDough, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.PrepareDough) * time.Second)

	assert.Equal(t, models.AddingIngredients, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.AddIngredients) * time.Second)

	assert.Equal(t, models.Baking, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.BakePizza) * time.Second)

	assert.Equal(t, models.Ready, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID))
}

func TestPizzaSpecificOrderStatus(t *testing.T) {
	testOrder := models.PizzaOrder{
		CustomerID: "SomeID2",
		PizzaType:  "VegPizza",
		PizzaSize:  "Medium",
		OrderID:    "SomeID2",
	}

	go CustomerOrder(testOrder)
	time.Sleep(time.Second * 1)

	assert.Equal(t, models.PreparingDough.String(), GetSpecificOrderInfo(testOrder.CustomerID, testOrder.OrderID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.PrepareDough) * time.Second)

	assert.Equal(t, models.AddingIngredients.String(), GetSpecificOrderInfo(testOrder.CustomerID, testOrder.OrderID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.AddIngredients) * time.Second)

	assert.Equal(t, models.Baking.String(), GetSpecificOrderInfo(testOrder.CustomerID, testOrder.OrderID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.BakePizza) * time.Second)

	assert.Equal(t, models.Ready.String(), GetSpecificOrderInfo(testOrder.CustomerID, testOrder.OrderID))

	assert.Equal(t, "", GetSpecificOrderInfo("non-exisitent", testOrder.OrderID))
	assert.Equal(t, models.NotFound.String(), GetSpecificOrderInfo(testOrder.CustomerID, "non-exisitent"))

}

func TestPizzaAllOrderStatus(t *testing.T) {
	testOrder := models.PizzaOrder{
		CustomerID: "SomeID2",
		PizzaType:  "VegPizza",
		PizzaSize:  "Medium",
		OrderID:    "SomeID2",
	}

	go CustomerOrder(testOrder)
	time.Sleep(time.Second * 1)

	allOrderStatus := []models.CustomerAllOrderStatus{
		models.CustomerAllOrderStatus{OrderID: "SomeID2"},
	}

	allOrderStatus[0].OrderStatus = models.PreparingDough.String()
	assert.Equal(t, allOrderStatus, GetAllOrderStatus(testOrder.CustomerID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.PrepareDough) * time.Second)

	allOrderStatus[0].OrderStatus = models.AddingIngredients.String()
	assert.Equal(t, allOrderStatus, GetAllOrderStatus(testOrder.CustomerID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.AddIngredients) * time.Second)

	allOrderStatus[0].OrderStatus = models.Baking.String()
	assert.Equal(t, allOrderStatus, GetAllOrderStatus(testOrder.CustomerID))
	time.Sleep(time.Duration(models.PizzaPreparationTime.BakePizza) * time.Second)

	allOrderStatus[0].OrderStatus = models.Ready.String()
	assert.Equal(t, allOrderStatus, GetAllOrderStatus(testOrder.CustomerID))

	assert.Equal(t, []models.CustomerAllOrderStatus(nil), GetAllOrderStatus("non-exisitent"))

}
