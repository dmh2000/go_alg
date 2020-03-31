package algsort

// InsertionSort ...
func InsertionSort(data AlgSort) {
	if data.length() == 0 {
		return
	}

	n := data.length()

	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.lt(j, j-1); j-- {
			data.exch(j, j-1)
		}
	}
}

// InsertionSort ...
func insertionSort(data AlgSort, lo int, hi int) {
	if lo >= hi {
		return
	}

	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && data.lt(j, j-1); j-- {
			data.exch(j, j-1)
		}
	}
}
