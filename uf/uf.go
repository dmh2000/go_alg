package uf

// Pair : input data
type Pair struct {
	P int
	Q int
}

// UF interface for union-find
// Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition.
type UF interface {
	// UF : initialize with a slice of pairs
	Uf(n int)

	// Union : adds connection between p and q
	Union(p int, q int)

	// Find : component index for p (0..N-1)
	Find(p int) int

	// Connected: the invariant is that two elements
	// are connected if and only they are in the same
	// component. They both have the same component id
	Connected(p int, q int) bool

	// Count : number of components
	Count() int
}
