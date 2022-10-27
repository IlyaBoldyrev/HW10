package main

import (
	"bufio"
	"fmt"
	"github.com/IlyaBoldyrev/HW10/sliceSort/sortFunctions"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Введите значения:")
	var inputSlice []int64 = make([]int64, 0, 25)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		inputSlice = append(inputSlice, num)
	}
	var inputSlice2 []int64 = make([]int64, len(inputSlice))
	copy(inputSlice2, inputSlice)
	sortFunctions.SortByMe(inputSlice)
	fmt.Println(inputSlice)
	sortFunctions.BuiltInSorting(inputSlice2)
	fmt.Println(inputSlice2)
}
