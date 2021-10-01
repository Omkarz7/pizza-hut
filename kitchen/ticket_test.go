package kitchen

import (
	"testing"
	"time"

	"github.com/Omkarz7/pizza-hut/models"
	"github.com/stretchr/testify/assert"
)

func TestShortURLGenerator(t *testing.T) {
	testOrder := models.PizzaOrder{
		CustomerID: "SomeID",
		PizzaType:  "VegPizza",
		PizzaSize:  "Medium",
	}

	go CustomerOrder(testOrder)

	assert.Equal(t, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID), models.Recieved)
	time.Sleep(time.Second * 1)

	assert.Equal(t, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID), models.PreparingDough)
	time.Sleep(time.Duration(models.PizzaPreparationTime.PrepareDough) * time.Second)

	assert.Equal(t, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID), models.AddingIngredients)
	time.Sleep(time.Duration(models.PizzaPreparationTime.AddIngredients) * time.Second)

	assert.Equal(t, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID), models.Baking)
	time.Sleep(time.Duration(models.PizzaPreparationTime.BakePizza) * time.Second)

	assert.Equal(t, models.CustomerList.GetStatus(testOrder.CustomerID, testOrder.OrderID), models.Ready)
}
