package algsort

// InsertionSort ...
func InsertionSort(data AlgSort) {
	n := data.length()

	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.lt(j, j-1); j-- {
			data.exch(j, j-1)
		}
	}
}
