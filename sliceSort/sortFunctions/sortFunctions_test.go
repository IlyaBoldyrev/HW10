package sortFunctions

import "testing"

var sl = []int64{3, 67, 45, 4335, 5354554, 45, 67, 0, -15, -3453, 67, -67, 0, 3434}

func BenchmarkSortByMe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortByMe(sl)
	}
}

func BenchmarkBuiltInSorting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuiltInSorting(sl)
	}
}
