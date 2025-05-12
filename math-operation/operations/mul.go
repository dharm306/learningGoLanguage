package operations

func Mul(no1 int, no2 int) (int, string) {
	if no1 <= 0 || no2 <= 0 {
		return 0, "Both number should be greater than Zero"
	}
	return no1 * no2, ""
}
