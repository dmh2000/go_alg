package algsort

func partition(data AlgSort, lo int, hi int) int {
	i := lo
	j := hi + 1
	for {
		for {
			i++
			if !data.lt(i, lo) {
				break
			}
			if i == hi {
				break
			}
		}
		for {
			j--
			if !data.lt(lo, j) {
				break
			}
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		data.exch(i, j)
	}
	data.exch(lo, j)
	return j
}

func sort(data AlgSort, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(data, lo, hi)
	sort(data, lo, j-1)
	sort(data, j+1, hi)
}

// QuickSort ...
func QuickSort(data AlgSort) {
	sort(data, 0, data.length()-1)
}
