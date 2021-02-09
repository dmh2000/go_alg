package algsort

// AlgSort interface definition
// implements sort.Interface
type AlgSort interface {
	isSorted() bool

	get(a int) int

	set(index int, value int)

	clone() AlgSort

	Len() int

	Swap(i,j int)

	Less(i,j int) bool
}
