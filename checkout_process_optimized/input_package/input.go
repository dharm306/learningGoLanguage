package inputpackage

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// shared prompt logic
func printPrompt(message string, val ...any) {
	if val == nil {
		fmt.Print(message)
	} else {
		fmt.Printf(message, val...)
	}
}

func GetIntValue(message string, val ...any) int {
	printPrompt(message, val...)

	var input int
	fmt.Scan(&input)
	return input
}

func GetFloatValue(message string, val ...any) float64 {
	printPrompt(message, val...)

	var input float64
	fmt.Scan(&input)
	return input
}

func GetStringValue(message string, val ...any) string {
	printPrompt(message, val...)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
