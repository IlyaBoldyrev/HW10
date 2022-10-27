package factorial

func Factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		num *= Factorial(num - 1)
		return num
	}
}
