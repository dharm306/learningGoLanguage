package order

import (
	"cp/iopackage"
	"fmt"
)

type product_type struct {
	id    int
	name  string
	price float64
	qty   int
}

type order_type struct {
	BillNo        string
	BillDate      string
	CustomerName  string
	OrderProducts []product_type
	FinalPrice    float64
}

// Static product list
var ProductList = []product_type{
	{id: 1, name: "Laptop", price: 1200.00, qty: 10},
	{id: 2, name: "Smartphone", price: 800.00, qty: 20},
	{id: 3, name: "Headphones", price: 150.00, qty: 50},
	{id: 4, name: "Speaker", price: 150.00, qty: 5},
}

var DraftOrder *order_type

func ProcessOrder(draftOrder *order_type, action int, data any) *order_type {
	switch action {
	case 1:
		ShowProductList(ProductList)
		selectedItem, isValid := getProductByID(iopackage.GetIntValue("Select Product : ", ""))
		if isValid {
			selectedQty := iopackage.GetIntValue("Select Qty of %s : ", selectedItem.name)
			if selectedQty <= selectedItem.qty {
				selectedItem.qty = selectedQty
				draftOrder.OrderProducts = append(draftOrder.OrderProducts, selectedItem)
			} else {
				fmt.Println("Please add product again with proper quantity.")
				return ProcessOrder(draftOrder, 1, data)
			}
		} else {
			fmt.Println("Invalid Product selected, Please select again.")
			return ProcessOrder(draftOrder, action, data)
		}
	case 3:
		productTotalPrice := 0.0
		for _, p := range draftOrder.OrderProducts {
			productTotalPrice += p.price * float64(p.qty)
		}
		draftOrder.FinalPrice = productTotalPrice
		return draftOrder
	}
	fmt.Println("1 Add Item to cart")
	fmt.Println("2 Show Total")
	fmt.Println("3 Complete Order")
	action = iopackage.GetIntValue("Please select Action : ", "")
	return ProcessOrder(draftOrder, action, data)
}

func ShowProductList(products []product_type) {
	fmt.Println("************* Product List ****************")
	fmt.Printf("%-2s %-15s %-10s %2s\n", "ID", "Product", "Price", "Quantity")
	for _, product := range products {
		fmt.Printf("%-2d %-15s %-10.2f %2d\n", product.id, product.name, product.price, product.qty)
	}
	fmt.Println("***********************************")
}

func getProductByID(id int) (product_type, bool) {
	for _, p := range ProductList {
		if p.id == id {
			return p, true // Found
		}
	}
	return product_type{}, false // Not found
}

func CreateDraftOrder(name, billNo, billDate string) *order_type {
	return &order_type{
		BillNo:        billNo,
		BillDate:      billDate,
		CustomerName:  name,
		OrderProducts: []product_type{},
		FinalPrice:    0.0,
	}
}
