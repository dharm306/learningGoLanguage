package main

import (
	"fmt"

	"calc.com/utility"
)

func main() {
	var StudentName string = "dharmesh patel"
	var Standard int = 10
	fmt.Println("Student Name is : ", utility.UcFirst(StudentName))
	fmt.Println("Standard is : ", Standard)
	consumedMarks, totalMarks := 0, 0
	fmt.Print("Please enter total Marks : ")
	fmt.Scan(&totalMarks)
	fmt.Printf("Please enter consumed Marks from %d : ", totalMarks)
	fmt.Scan(&consumedMarks)
	var percentage float64 = (float64(consumedMarks) / float64(totalMarks)) * 100

	fmt.Printf("%s has obtaind %s percentage in %d Standard \n", utility.UcFirst(StudentName), utility.FormatNumber(percentage), Standard)

}
