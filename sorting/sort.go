package algsort

// AlgSort interface definition
type AlgSort interface {
	isSorted() bool

	exch(a int, b int)

	lt(a int, b int) bool

	length() int

	get(a int) int

	set(index int, value int)

	clone() AlgSort
}
