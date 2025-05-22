package main

import (
	"fmt"
	"s6/iopackage"
)

func main() {
	fmt.Println("This is session 6")
	fmt.Println("We have covered slices : advcance array")
	var subjectNames []string
	var marks []int
	var totalSubject int = iopackage.GetIntValue("Enter total required subject :", "")
	for i := range totalSubject {
		subjectNames = append(subjectNames, iopackage.GetStringValue("Please name of Subject %d : ", i+1))
	}
	for j := range totalSubject {
		marks = append(marks, iopackage.GetIntValue("Please enter mark of %s : ", subjectNames[j]))
	}

	fmt.Println("--------------------------------------")
	for k := range totalSubject {
		//fmt.Println("Subject ", subjectNames[k])
		//fmt.Println("Marks ", marks[k])
		iopackage.DisplayString("%-20s %3d\n", subjectNames[k], marks[k])
	}
	fmt.Println("--------------------------------------")

}
