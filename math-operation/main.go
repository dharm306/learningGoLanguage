package main

import (
	"bufio"
	"fmt"
	"os"

	"math.com/operations"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	no1, no2 := 0, 0
	fmt.Println("Welcome to the Math Operation Program")
	fmt.Print("Please enter your name : ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", name)
	fmt.Print("Please enter first no :")
	fmt.Scan(&no1)
	fmt.Print("Please enter second no :")
	fmt.Scan(&no2)

	fmt.Printf("Your selected no are %d,%d.\n", no1, no2)

	action := 0
	fmt.Println("1 - Add")
	fmt.Println("2 - Sub")
	fmt.Println("3 - Div")
	fmt.Println("4 - Mul")

	fmt.Print("Please enter no for what operation do you want to perform ? ")
	fmt.Scan(&action)
	fmt.Printf("Your selected action no is %d.\n", action)
	answer, error := 0, ""
	if action == 1 {
		answer, error = operations.Sum(no1, no2)
	} else if action == 2 {
		answer, error = operations.Sub(no1, no2)
	} else if action == 3 {
		answer, error = operations.Div(no1, no2)
	} else if action == 4 {
		answer, error = operations.Mul(no1, no2)
	} else {
		error = "Invalid Operation selected."
	}
	if error != "" {
		fmt.Printf("There is an error in Operation : %s\n", error)
	} else {
		fmt.Printf("Answer of action no %d is %d \n", action, answer)
	}
}
