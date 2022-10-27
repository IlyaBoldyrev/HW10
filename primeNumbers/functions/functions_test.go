package functions

import "fmt"

func ExempleGetPrimeNumbers() {
	sl := GetPrimeNumbers(14)
	fmt.Println(sl)
}

func ExemplePrintSliceOfInts() {
	sl := []int{1, 10, 13, 8, 5, 4}
	PrintSliceOfInts(sl)
}
