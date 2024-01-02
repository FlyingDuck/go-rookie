package util

func MaxInt(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func MaxIntEnhanced(num1, num2 int) (max int, first bool) {
	if num1 > num2 {
		return num1, true
	}
	return num2, false
}
