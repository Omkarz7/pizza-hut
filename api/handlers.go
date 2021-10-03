package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Omkarz7/pizza-hut/kitchen"
	"github.com/Omkarz7/pizza-hut/models"
	"github.com/google/uuid"
)

// BuyPizza handle buy_pizza route to order pizza
func BuyPizza(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var customerOrder models.PizzaOrder
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &customerOrder)
		customerOrder.OrderID = uuid.New().String()
		go kitchen.CustomerOrder(customerOrder)
		fmt.Fprintf(w, "You order has been placed! Your Order ID is %s", customerOrder.OrderID)
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}

// OrderStatus retrieves order status for a certain pizza
func OrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		customerID := r.URL.Query().Get("customerID")
		orderID := r.URL.Query().Get("orderID")

		if customerID == "" {
			http.Error(w, "Customer ID missing. Failed to retrieve data", http.StatusNotFound)
		} else if customerID != "" && orderID != "" {
			orderInfoList := kitchen.GetSpecificOrderInfo(customerID, orderID)
			if orderInfoList == "" {
				http.Error(w, "No order found for the given customerID", http.StatusNotFound)
				return
			}
			fmt.Fprintf(w, "Order status:%s", orderInfoList)
		} else if customerID != "" {
			//retireve all order
			orderInfoList := kitchen.GetAllOrderStatus(customerID)
			if orderInfoList == nil {
				http.Error(w, "No order found for the given customerID", http.StatusNotFound)
				return
			}
			fmt.Fprintf(w, "Order status:"+fmt.Sprintf("%+v", orderInfoList))
		}
	} else {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
	}
}

// HomePagePing is a dummy URL test if server is up
func HomePagePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Pizzaz!")
}
