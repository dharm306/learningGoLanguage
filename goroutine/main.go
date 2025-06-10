package main

import (
	"fmt"
	"time"
)

func showValue(v int) {
	fmt.Println("Number is : ", v)
}

func main() {
	fmt.Println("This program is for clear concept of Goroutine")

	for i := range 10 {
		go showValue(i)
	}

	time.Sleep(time.Second)
}
