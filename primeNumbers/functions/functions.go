package functions

import "fmt"

// GetPrimeNumbers calculate all prime numbers up to n
func GetPrimeNumbers(n float64) []int {
	sl := make([]int, 0, 1000)
	for i := 2; i <= int(n); i++ {
		k := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				k++
			}
		}
		if k == 2 {
			sl = append(sl, i)
		}
	}
	return sl
}

// PrintPrimeNumbers print slice of ints
func PrintSliceOfInts(sl []int) {
	for _, n := range sl {
		fmt.Println(n)
	}
}
