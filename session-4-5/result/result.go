package result

import "fmt"

var TotalMarks float64 = 600
var Subjects = [6]string{"Maths", "Science", "Gujarati", "English", "Hindi", "SS"}

func CalculatePercentage(gainedMarks int) float64 {
	return (convertIntToFloat(gainedMarks) / TotalMarks) * 100
}

func convertIntToFloat(value int) float64 {
	return float64(value)
}

func DisplayMarks(gainedMarks [6]int) {
	println("------------ Marks --------")
	for i := 0; i < len(gainedMarks); i++ {
		fmt.Printf("%s\t\t\t:%d\n", Subjects[i], gainedMarks[i])
	}
	println("---------------------------")
}
