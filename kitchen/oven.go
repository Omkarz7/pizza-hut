package kitchen

import (
	"time"

	"github.com/Omkarz7/pizza-hut/models"
)

func bakePizza() {
	time.Sleep(time.Duration(models.PizzaPreparationTime.BakePizza) * time.Second)
}
