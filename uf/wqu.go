package uf

/*
implementation of Weighted-Quick-Union algorithm.
Analysis of Weighted-Quick-Union is presented in
Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter One
*/

// Wqu : weighted-quick union implementation
type Wqu struct {
	id    []int // list of components
	sz    []int // size of component trees
	count int
}

// NewWeightedQuickUnion : create and init a quick-union object
func NewWeightedQuickUnion(n int) *Wqu {
	wqu := Wqu{}
	wqu.Uf(n)
	return &wqu
}

// Uf : initialize with a slice of pairs
func (wqu *Wqu) Uf(n int) {
	// store the current count
	wqu.count = n

	// initilaize the list of components
	wqu.id = make([]int, n)
	for i := range wqu.id {
		wqu.id[i] = i
	}

	// initialize tree size
	wqu.sz = make([]int, n)
	for i := range wqu.sz {
		wqu.sz[i] = 1
	}
}

// Union : adds connection between p and q
func (wqu *Wqu) Union(p int, q int) {
	// get the current components of p and q
	pid := wqu.Find(p)
	qid := wqu.Find(q)

	// if they match, nothing more to do
	if pid == qid {
		return
	}

	if pid == qid {
		return
	}

	// attach smaller tree to larger
	if wqu.sz[pid] < wqu.sz[qid] {
		// add q to tree of p
		wqu.id[pid] = q
		wqu.sz[qid] += wqu.sz[pid]
	} else {
		// add p to tree of q
		wqu.id[qid] = p
		wqu.sz[pid] += wqu.sz[qid]
	}

	// reduce live count by 1 because p became q
	wqu.count--
}

// Find : component index for p (0..N-1)
func (wqu *Wqu) Find(p int) int {
	// p is in component id[p]
	// follow the tree branch to find p
	for p != wqu.id[p] {
		p = wqu.id[p]
	}
	return p
}

// Connected : return true if p and q are connected
func (wqu *Wqu) Connected(p int, q int) bool {
	// connected only if both p and q are in the same component
	return wqu.Find(p) == wqu.Find(q)
}

// Count : number of components
func (wqu *Wqu) Count() int {
	return wqu.count
}
