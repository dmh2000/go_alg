package algsort

func printM(data AlgSort, lo int, hi int) {
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

func merge(data AlgSort, aux AlgSort, lo int, mid int, hi int) {
	i := lo
	j := mid + 1
	k := lo
outer:
	for {
		// no more in left half
		if i == mid+1 {
			for j <= hi {
				aux.set(k, data.get(j))
				j++
				k++
			}
			break outer
		}
		// no more in right half
		if j == hi+1 {
			for i <= mid {
				aux.set(k, data.get(i))
				i++
				k++
			}
			break outer
		}
		if k > hi {
			break outer
		}

		if data.Less(i, j) {
			aux.set(k, data.get(i))
			i++
			k++
		} else {
			aux.set(k, data.get(j))
			j++
			k++
		}
	}
	// copy all back to data
	for i := lo; i <= hi; i++ {
		data.set(i, aux.get(i))
	}
}

func msort(data AlgSort, aux AlgSort, lo int, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2

	msort(data, aux, lo, mid)
	msort(data, aux, mid+1, hi)
	merge(data, aux, lo, mid, hi)
}

// MergeSort ...
func MergeSort(data AlgSort) {
	if data.Len() == 0 {
		return
	}
	aux := data.clone()

	msort(data, aux, 0, data.Len()-1)
}
