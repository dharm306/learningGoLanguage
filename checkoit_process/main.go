package main

import (
	"cp/iopackage"
	"fmt"
	"time"
)

type product struct {
	id    int
	name  string
	price float64
	qty   int
}

type order struct {
	billNo        string
	billDate      string
	customerName  string
	orderProducts []product
	finalPrice    float64
}

// Static product list
var productList = []product{
	{id: 1, name: "Laptop", price: 1200.00, qty: 10},
	{id: 2, name: "Smartphone", price: 800.00, qty: 20},
	{id: 3, name: "Headphones", price: 150.00, qty: 50},
	{id: 4, name: "Speaker", price: 150.00, qty: 5},
}

func main() {
	fmt.Println("Welcome to the Checkout Process")
	now := time.Now()
	draftOrder := order{
		billNo:        now.Format("20060102150405"),
		billDate:      now.Format("02-01-2006"),
		customerName:  iopackage.GetStringValue("Enter Customer Name : ", ""),
		orderProducts: []product{},
		finalPrice:    0.0}

	processOrder(&draftOrder, 0, nil)

	fmt.Println("*************** Order Info *****************")
	fmt.Printf("Customer Name : %s\n", draftOrder.customerName)
	fmt.Printf("Bill No : %-10s Bill Data : %10s\n", draftOrder.billNo, draftOrder.billDate)
	showProductList(draftOrder.orderProducts)
	fmt.Printf("Order Total : %20.2f\n", draftOrder.finalPrice)
}

func processOrder(draftOrder *order, action int, data any) *order {
	switch action {
	case 1:
		showProductList(productList)
		selectedItem, isValid := getProductByID(iopackage.GetIntValue("Select Product : ", ""))
		if isValid {
			selectedQty := iopackage.GetIntValue("Select Qty of %s : ", selectedItem.name)
			if selectedQty <= selectedItem.qty {
				selectedItem.qty = selectedQty
				draftOrder.orderProducts = append(draftOrder.orderProducts, selectedItem)
			} else {
				fmt.Println("Please add product again with proper quantity.")
				return processOrder(draftOrder, 1, data)
			}
		} else {
			fmt.Println("Invalid Product selected, Please select again.")
			return processOrder(draftOrder, action, data)
		}
	case 3:
		productTotalPrice := 0.0
		for _, p := range draftOrder.orderProducts {
			productTotalPrice += p.price * float64(p.qty)
		}
		draftOrder.finalPrice = productTotalPrice
		return draftOrder
	}
	fmt.Println("1 Add Item to cart")
	fmt.Println("2 Show Total")
	fmt.Println("3 Complete Order")
	action = iopackage.GetIntValue("Please select Action : ", "")
	return processOrder(draftOrder, action, data)
}

func showProductList(products []product) {
	fmt.Println("************* Product List ****************")
	fmt.Printf("%-2s %-15s %-10s %2s\n", "ID", "Product", "Price", "Quantity")
	for _, product := range products {
		fmt.Printf("%-2d %-15s %-10.2f %2d\n", product.id, product.name, product.price, product.qty)
	}
	fmt.Println("***********************************")
}

func getProductByID(id int) (product, bool) {
	for _, p := range productList {
		if p.id == id {
			return p, true // Found
		}
	}
	return product{}, false // Not found
}
