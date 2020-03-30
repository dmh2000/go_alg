package algsort

// Sort interface definition
type AlgSort interface {
	isSorted() bool

	exch(a int, b int)

	lt(a int, b int) bool

	length() int
}
