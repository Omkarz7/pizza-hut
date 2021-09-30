package notify

import (
	"os"

	"github.com/Omkarz7/pizza-hut/models"
)

func NotifyCustomer() {

	var orderStatus string
	f, err := os.OpenFile("CompletedOrders.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for {
		select {
		case orderStatus = <-models.NotifyStatus:
			if _, err = f.WriteString(orderStatus); err != nil {
				panic(err)
			}
		}
	}

}
