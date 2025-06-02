package main

import (
	inputpackage "cpo/input_package"
	"cpo/order"
	"cpo/product"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("This is optimized checkout process.")
	now := time.Now()
	//customerName := inputpackage.GetStringValue("Please enter your name: ", nil)
	//println("Customer Name:", customerName)
	//product.ListAllProducts()

	//selectedProduct := (product.GetProductByID(1))

	// order.InitOrder()
	// order.DraftOrder.CustomerName = "John Doe"
	// order.DraftOrder.BillNo = "BILL12345"
	// order.DraftOrder.BillDate = "2023-10-01"
	// order.DraftOrder.AddProduct(selectedProduct, 5)
	// order.StoreOrder(order.DraftOrder)
	// product.ListAllProducts()
	choice := 0
	for orderAction := true; orderAction; {
		switch choice {

		case 1:
			order.InitOrder()
			order.DraftOrder.CustomerName = inputpackage.GetStringValue("Please enter customer name: ")
			order.DraftOrder.BillNo = now.Format("02012006") + strconv.Itoa(len(*order.AllOrders)+1)
			order.DraftOrder.BillDate = now.Format("02-01-2006")
			choice = 2
		case 2:
			if order.DraftOrder != nil {

				product.ListAllProducts()
				productID := inputpackage.GetIntValue("Enter Product ID to add from from %d to %d: ", 1, len(product.ProductList))

				selectedProduct, isFound := product.GetProductByID(productID)
				if isFound {
					quantity := inputpackage.GetIntValue("Enter Quantity (less than %d): ", selectedProduct.Quantity)
					order.DraftOrder.AddProduct(selectedProduct, quantity)
					fmt.Printf("Added %d of %s to the order.\n", quantity, selectedProduct.Name)

					if addMore := strings.ToLower(inputpackage.GetStringValue("Do you want to add more products? (yes/no): ")); addMore == "no" {
						choice = 3
						fmt.Println("New Choise is ", choice)
					}
				} else {
					fmt.Println("Invalid Product ID.")
				}
			} else {
				fmt.Println("No draft order found. Please create an order first.")
				choice = 1
			}
		case 3:
			order.DraftOrder.CompleteOrder()
			order.ResetOrder()
			fmt.Println("Order completed successfully!")
			if addMore := strings.ToLower(inputpackage.GetStringValue("Do you want create new order? (yes/no): ")); addMore == "no" {
				choice = 4
			} else {
				choice = 1
			}
		case 4:
			order.ListAllOrders()
			fmt.Println("Exiting the checkout process.")
			orderAction = false
		default:
			fmt.Println("1. Create Order")
			fmt.Println("2. Add Product to Order")
			fmt.Println("3. Complete Order")
			fmt.Println("4. Exit")
			choice = inputpackage.GetIntValue("Please select an option from %d to %d : ", 1, 4)

		}

	}
	//fmt.Printf("Draft Order: %+v\n", order.DraftOrder)
	//fmt.Printf("Draft Order: %+v\n", order.AllOrders)
}
