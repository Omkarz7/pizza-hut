package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Omkarz7/pizza-hut/api"
	"github.com/Omkarz7/pizza-hut/models"
	"github.com/Omkarz7/pizza-hut/notify"
)

func main() {
	go notify.NotifyCustomer()
	fmt.Println("Starting server at port ", models.Port)
	http.HandleFunc("/", api.HomePagePing)
	http.HandleFunc("/buy_pizza", api.BuyPizza)
	http.HandleFunc("/order_status", api.OrderStatus)

	if err := http.ListenAndServe(models.Port, nil); err != nil {
		log.Fatal(err)
	}
}
