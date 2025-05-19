package main

import (
	"fmt"
	"s4/iopackage"
	"s4/result"
	typeconversation "s4/type-conversation"
)

func main() {
	fmt.Println("This is session 4")
	var marks [6]int
	totalGainedMarks := 0
	stringVal := "25.50"
	floatVal := 240.50
	fmt.Printf("%s String to Integer : %d \n", stringVal, typeconversation.StringToInt(stringVal))
	fmt.Printf("%.2f Float to Integer : %d \n", floatVal, typeconversation.FloatToInt(floatVal))
	fmt.Println("----------------------------------------------------------")
	rollNo := iopackage.GetIntValue("Please enter Student Roll NO :", "")
	iopackage.DisplayInteger("Your entered Roll No is %d\n", rollNo)
	for i := 0; i < len(result.Subjects); i++ {
		marks[i] = iopackage.GetIntValue("Please enter %s Mark :", result.Subjects[i])
		totalGainedMarks += marks[i]
	}
	percentage := result.CalculatePercentage(totalGainedMarks)
	result.DisplayMarks(marks)
	fmt.Printf("Calculated Percentage : %.2f\n", percentage)
}
