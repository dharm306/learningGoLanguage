package operations

func Sub(no1 int, no2 int) (int, string) {
	if no2 > no1 {
		return 0, "No2 should not be greater than No1"
	}

	return no1 - no2, ""
}
