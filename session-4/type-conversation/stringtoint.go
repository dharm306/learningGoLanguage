package typeconversation

import "strconv"

func StringToInt(input string) int {
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0 // or handle the error as needed
	}
	return int(val) // truncates decimal part
}
