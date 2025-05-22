package main

import (
	"fmt"
	"s7/iopackage"
)

func main() {
	totalMarks := 300
	totalStudent := iopackage.GetIntValue("Enter Number Of Student : ", "")
	// var testData = make(map[string]int)
	// testData["Dharmesh"] = 35
	// testData["Rakesh"] = 35

	// fmt.Println("TestData", testData)

	// delete(testData, "Rakesh")

	// fmt.Println("TestData", testData)

	var studentInfo = make(map[int]map[string]int)

	for i := range totalStudent {

		fmt.Printf("Enter Student %d Marks \n", i+1)
		studentInfo[i] = make(map[string]int)
		studentInfo[i]["English"] = iopackage.GetIntValue("Enter English Subject Mark : ", "")
		studentInfo[i]["Science"] = iopackage.GetIntValue("Enter Science Subject Mark : ", "")
		studentInfo[i]["Gujarati"] = iopackage.GetIntValue("Enter Science Subject Mark : ", "")
		fmt.Println("-------------------------")
	}

	fmt.Println("studentInfo", studentInfo)

	for j := range totalStudent {
		fmt.Println("------------------------------")
		fmt.Printf("Student %d Result \n", j+1)
		totalObtinedMarks := 0
		fmt.Println("------------------------------")
		for subject, value := range studentInfo[j] {
			//fmt.Println("studentInfo", subject, value)
			fmt.Printf("%-20s %d\n", subject, value)
			totalObtinedMarks += int(value)
		}
		fmt.Println("------------------------------")
		percentage := (float64(totalObtinedMarks) / float64(totalMarks)) * 100
		fmt.Printf("Obtained Marks : %d / %d\n", totalObtinedMarks, totalMarks)
		fmt.Println("------------------------------")
		fmt.Printf("Percentage : %.2f\n", percentage)
		fmt.Println("###############################")
	}
}
