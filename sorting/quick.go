package algsort

import "math/rand"

func printQ(data AlgSort, lo int, hi int) {
	// fmt.Printf("%v : %v : ", lo, hi)
	// for i := 0; i < lo; i++ {
	// 	fmt.Printf("%v ", data.at(i))
	// }
	// fmt.Printf("|")
	// for i := lo; i <= hi; i++ {
	// 	fmt.Printf("%v ", data.at(i))
	// }
	// fmt.Printf("|")
	// for i := hi + 1; i < data.length(); i++ {
	// 	fmt.Printf("%v ", data.at(i))
	// }
	// fmt.Printf("\n")
}

// shuffle : fisher-yates shuffle
func shuffle(data AlgSort) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		j := rand.Int() % (n - i)
		j += i
		data.Swap(i, j)
	}
}

func partition(data AlgSort, lo int, hi int) int {
	i := lo
	j := hi + 1
	for {
		for {
			i++
			if !data.Less(i, lo) {
				break
			}
			if i == hi {
				break
			}
		}
		for {
			j--
			if !data.Less(lo, j) {
				break
			}
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		data.Swap(i, j)

	}
	data.Swap(lo, j)
	printQ(data, lo, hi)
	return j
}

func qsort(data AlgSort, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(data, lo, hi)
	qsort(data, lo, j-1)
	qsort(data, j+1, hi)
}

// QuickSort ...
func QuickSort(data AlgSort) {
	if data.Len() == 0 {
		return
	}

	// shuffle to avoid n^2 for inorder sort
	shuffle(data)
	// quicksort
	qsort(data, 0, data.Len()-1)
}
