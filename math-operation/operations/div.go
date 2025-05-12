package operations

func Div(no1 int, no2 int) (int, string) {
	if no2 <= 0 {
		return 0, "No2 should be greater than Zero"
	}

	return no1 / no2, ""
}
