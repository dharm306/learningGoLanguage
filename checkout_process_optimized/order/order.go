package order

import (
	"cpo/product"
	"fmt"
)

type TypeOrder struct {
	BillNo        string
	BillDate      string
	CustomerName  string
	OrderProducts []product.Product
	FinalPrice    float64
	status        string // draft, completed
}

var AllOrders *[]TypeOrder = &[]TypeOrder{}

var DraftOrder *TypeOrder

func InitOrder() {
	DraftOrder = &TypeOrder{
		BillNo:        "",
		BillDate:      "",
		CustomerName:  "",
		OrderProducts: []product.Product{},
		FinalPrice:    0.0,
		status:        "draft",
	}
}

func ResetOrder() {
	DraftOrder = nil
}

func (draftOrder *TypeOrder) AddProduct(p *product.Product, qty int) {
	if p.Quantity >= qty {
		p.Quantity -= qty
		draftOrder.OrderProducts = append(draftOrder.OrderProducts, product.Product{
			ID:       p.ID,
			Name:     p.Name,
			Price:    p.Price,
			Quantity: qty, // override with purchased quantity
		})
	} else {
		// Handle insufficient stock case
	}
}

func StoreOrder(o *TypeOrder) {
	*AllOrders = append(*AllOrders, *o)
}

func (do *TypeOrder) CompleteOrder() {
	do.status = "completed"
	do.FinalPrice = 0.0
	for _, p := range do.OrderProducts {
		do.FinalPrice += p.Price * float64(p.Quantity)
	}
	StoreOrder(do)
}

func ListAllOrders() {
	if len(*AllOrders) == 0 {
		println("No orders found.")
		return
	}
	fmt.Println("*********************** Order List ***********************")
	for _, order := range *AllOrders {
		fmt.Println("*******************************************************")
		fmt.Println("Bill No:", order.BillNo, "Customer Name:", order.CustomerName)
		fmt.Printf("Bill Date: %s, Status: %s\n", order.BillDate, order.status)
		fmt.Println("Products in Order:")
		fmt.Printf("%-3s\t%-15s %5s\t%-3s %5s\t\n", "ID", "Name", "Price", "Qty", "Total")
		for _, product := range order.OrderProducts {
			subTotal := product.Price * float64(product.Quantity)
			fmt.Printf("%-3d\t%-15s %5.2f\t%-3d %5.2f\t\n", product.ID, product.Name, product.Price, product.Quantity, subTotal)
		}
		fmt.Println("*******************************************************")
		fmt.Println("Total Price:", order.FinalPrice)
		fmt.Println("*******************************************************")
	}
}
