package uf

/*
implementation of Quick-Union algorithm.
Analysis of quick-Union is presented in
Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter one


*/

// Qu : quick union implementation
type Qu struct {
	id    []int // list of components
	count int
}

// NewQuickUnion : create and init a quick-union object
func NewQuickUnion(n int) Qu {
	qu := Qu{}
	qu.Uf(n)
	return qu
}

// Uf : initialize with a slice of pairs
func (qu *Qu) Uf(n int) {
	// store the current count
	qu.count = n
	// create a slice of size n
	qu.id = make([]int, n)
	// initilaize the list of components
	for i := range qu.id {
		qu.id[i] = i
	}
}

// Union : adds connection between p and q
func (qu *Qu) Union(p int, q int) {
	// get the current components of p and q
	pid := qu.Find(p)
	qid := qu.Find(q)

	// if they match, nothing more to do
	if pid == qid {
		return
	}

	// set p connected to q
	qu.id[pid] = qid

	// reduce live count by 1 because p became q
	qu.count--
}

// Find : component index for p (0..N-1)
func (qu *Qu) Find(p int) int {
	// p is in component id[p]
	// follow the tree branch to find p
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}

// Connected : return true if p and q are connected
func (qu *Qu) Connected(p int, q int) bool {
	// connected only if both p and q are in the same component
	return qu.Find(p) == qu.Find(q)
}

// Count : number of components
func (qu *Qu) Count() int {
	return qu.count
}
