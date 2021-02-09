package algsort

import (
	"sort"
)

// GoSort : builtin sort
func GoSort(data AlgSort) {
	sort.Sort(AlgSort(data))
}