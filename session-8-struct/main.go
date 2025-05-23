package main

import (
	"fmt"
	"s8/iopackage"
)

type typemarks struct {
	math int
	sci  int
	eng  int
}
type student struct {
	name  string
	std   string
	marks typemarks
}

func newStudent(n string, s string, m int, sc int, e int) *student {
	newStudent := student{
		name: n,
		std:  s,
		marks: typemarks{
			math: m,
			sci:  sc,
			eng:  e,
		},
	}

	return &newStudent
}

func main() {
	totalMarks := 300
	totalStudent := iopackage.GetIntValue("Enter Number Of Student : ", "")

	var studentInfo = make(map[int]*student)

	for i := range totalStudent {

		fmt.Printf("Enter Student %d Marks \n", i+1)

		name := iopackage.GetStringValue("Enter Student %d Name :", i+1)
		std := iopackage.GetStringValue("Enter Student %d Standard :", i+1)
		math := iopackage.GetIntValue("Enter Student %d Maths Mark :", i+1)
		sci := iopackage.GetIntValue("Enter Student %d Science Mark :", i+1)
		eng := iopackage.GetIntValue("Enter Student %d English Mark :", i+1)

		stud := newStudent(name, std, math, sci, eng)
		studentInfo[i] = stud

		fmt.Println("-------------------------")
	}

	for j := range totalStudent {
		fmt.Println("------------------------------")
		fmt.Printf("Student %d Result \n", j+1)
		totalObtinedMarks := 0
		fmt.Println("------------------------------")
		s := studentInfo[j]
		fmt.Printf("Student %d Name : %10s\n", j+1, s.name)
		fmt.Printf("Student %d Std : %10s\n", j+1, s.std)
		fmt.Println("----------------Marks--------------")
		fmt.Printf("%-20s %d\n", "Maths", s.marks.math)
		fmt.Printf("%-20s %d\n", "Science", s.marks.sci)
		fmt.Printf("%-20s %d\n", "English", s.marks.eng)
		totalObtinedMarks += int(s.marks.eng + s.marks.sci + s.marks.eng)
		fmt.Println("------------------------------")
		percentage := (float64(totalObtinedMarks) / float64(totalMarks)) * 100
		fmt.Printf("Obtained Marks : %d / %d\n", totalObtinedMarks, totalMarks)
		fmt.Println("------------------------------")
		fmt.Printf("Percentage : %.2f\n", percentage)
		fmt.Println("###############################")
	}
}
