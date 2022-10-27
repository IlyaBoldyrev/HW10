package main

import (
	"fmt"
	"github.com/IlyaBoldyrev/HW10/primeNumbers/functions"
	"math"
)

func main() {
	var n float64
	fmt.Println("Введите число N: ")
	fmt.Scan(&n)
	n = math.Abs(n)
	sl := functions.GetPrimeNumbers(n)
	functions.PrintSliceOfInts(sl)
}
