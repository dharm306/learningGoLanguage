package iopackage

import (
	"fmt"
	"reflect"
)

func GetStringValue(inputMessage string, value any) string {
	var val string
	v := reflect.ValueOf(value)
	if v.IsZero() {
		fmt.Print(inputMessage)
	} else {
		fmt.Printf(inputMessage, v)
	}
	fmt.Scan(&val)
	return val
}
func DisplayString(printMessage string, values ...any) {
	fmt.Printf(printMessage, values...)
}
