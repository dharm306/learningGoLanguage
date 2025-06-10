package main

import "fmt"

type OrderStatus string

const (
	OrderStatusDraft     OrderStatus = "Draft"
	OrderStatusPlaced    OrderStatus = "Placed"
	OrderStatusInProcess OrderStatus = "In Process"
)

func (s OrderStatus) String() string {
	return string(s)
}

func main() {
	fmt.Println("Draft Order ", OrderStatusDraft.String())
}
