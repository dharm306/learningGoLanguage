package main

import (
	"fmt"

	"dpdemo.com/utility"
)

func main() {
	fmt.Println("My Name Is ", utility.UcFirst("dharmesh patel"))
	fmt.Println("My Salary is  ", utility.FormatNumber(25000.5000))
}
