package notify

import (
	"os"

	"github.com/Omkarz7/pizza-hut/models"
)

func NotifyCustomer() {
	for orderStatus := range models.NotifyStatus {
		func() {
			f, err := os.OpenFile("CompletedOrders.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			if _, err = f.WriteString(orderStatus); err != nil {
				panic(err)
			}
		}()
	}
}
