package kitchen

import (
	"time"

	"github.com/Omkarz7/pizza-hut/models"
)

func prepareDough() {
	time.Sleep(time.Duration(models.PizzaPreparationTime.PrepareDough) * time.Second)
}

func addIngredients() {
	time.Sleep(time.Duration(models.PizzaPreparationTime.AddIngredients) * time.Second)
}
