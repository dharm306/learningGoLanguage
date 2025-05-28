package main

import (
	"cp/iopackage"
	"cp/order"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the Checkout Process")
	now := time.Now()

	// Assign return value to the package-level `DraftOrder`
	order.DraftOrder = order.CreateDraftOrder(
		iopackage.GetStringValue("Enter Customer Name : ", ""),
		now.Format("20060102150405"),
		now.Format("02-01-2006"),
	)
	order.ProcessOrder(order.DraftOrder, 0, nil)
	fmt.Println("*************** Order Info *****************")
	fmt.Printf("Customer Name : %s\n", order.DraftOrder.CustomerName)
	fmt.Printf("Bill No : %-10s Bill Data : %10s\n", order.DraftOrder.BillNo, order.DraftOrder.BillDate)
	order.ShowProductList(order.DraftOrder.OrderProducts)
	fmt.Printf("Order Total : %20.2f\n", order.DraftOrder.FinalPrice)
}
