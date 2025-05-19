package iopackage

import (
	"fmt"
	"reflect"
)

func GetIntValue(inputMessage string, value any) int {
	var val int
	v := reflect.ValueOf(value)
	if v.IsZero() {
		fmt.Print(inputMessage)
	} else {
		fmt.Printf(inputMessage, v)
	}
	fmt.Scan(&val)
	return val
}
func DisplayInteger(printMessage string, value int) {
	fmt.Printf(printMessage, value)
}
