package uf

/*
implementation of Quick-Find algorithm.
Analysis of quick-find is presented in
Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter one

*/

// Qf : quick find implementation
type Qf struct {
	id    []int // list of components
	count int
}

// NewQuickFind : create and init a Union-Find object
func NewQuickFind(n int) *Qf {
	qf := Qf{}
	qf.Uf(n)
	return &qf
}

// Uf : initialize with a slice of pairs
func (qf *Qf) Uf(n int) {
	// store the current count
	qf.count = n
	// create a slice of size n
	qf.id = make([]int, n)
	// initilaize the list of components
	for i := range qf.id {
		qf.id[i] = i
	}
}

// Union : adds connection between p and q
func (qf *Qf) Union(p int, q int) {
	// get the current components of p and q
	pid := qf.Find(p)
	qid := qf.Find(q)

	// if they match, nothing more to do
	if pid == qid {
		return
	}

	// change values from id[p] to id[q]
	for i := range qf.id {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
	// reduce live count by 1 because p became q
	qf.count--
}

// Find : component index for p (0..N-1)
func (qf Qf) Find(p int) int {
	// p is in component id[p]
	return qf.id[p]
}

// Connected : return true if p and q are connected
func (qf Qf) Connected(p int, q int) bool {
	// connected only if both p and q are in the same component
	return qf.Find(p) == qf.Find(q)
}

// Count : number of components
func (qf Qf) Count() int {
	return qf.count
}
