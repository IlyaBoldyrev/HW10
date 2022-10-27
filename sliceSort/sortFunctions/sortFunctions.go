package sortFunctions

import "sort"

func SortByMe(sl []int64) {
	for i := 1; i < len(sl); i++ {
		temp := sl[i]
		j := i
		for j > 0 && sl[j-1] > temp {
			sl[j] = sl[j-1]
			j--
		}
		sl[j] = temp
	}
}

func BuiltInSorting(sl []int64) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
}
