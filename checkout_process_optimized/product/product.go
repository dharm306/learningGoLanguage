package product

import "fmt"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

var ProductList = []Product{
	{ID: 1, Name: "Laptop", Price: 1000.00, Quantity: 10},
	{ID: 2, Name: "Smartphone", Price: 500.00, Quantity: 20},
	{ID: 3, Name: "Tablet", Price: 300.00, Quantity: 15},
	{ID: 4, Name: "Smartwatch", Price: 200.00, Quantity: 25},
	{ID: 5, Name: "Headphones", Price: 100.00, Quantity: 30},
	{ID: 6, Name: "Bluetooth Speaker", Price: 150.00, Quantity: 18},
	{ID: 7, Name: "Camera", Price: 800.00, Quantity: 12}}

func GetProductByID(id int) (*Product, bool) {
	for _, product := range ProductList {
		if product.ID == id {
			return &product, true
		}
	}
	return &Product{}, false
}

func ListAllProducts() {
	fmt.Println("*************** Product List ***************")
	fmt.Printf("%-3s\t%-20s %5s\t%-3s\n", "ID", "Name", "Price", "Qty")
	for _, product := range ProductList {
		fmt.Printf("%-3d\t%-20s %5.2f\t%-3d\n", product.ID, product.Name, product.Price, product.Quantity)
	}
	fmt.Println("********************************************")
}
