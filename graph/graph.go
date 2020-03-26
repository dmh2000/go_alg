package graph

/*
	Go code in this file is based on the Java implementation from
	Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter 4
*/

// Graph :  interface to undirected graph
// this implemntation assumes the number of vertices are known at init time.
// edges can be added dynamically
// Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter 4
type Graph interface {
	// Graph : create a graph with no edges
	Graph(v int)

	// V : number of vertices
	V() int

	// E : number of edges
	E() int

	// AddEdge : add an edge from v to w
	AddEdge(v int, w int)

	//Adj : get a list of vertices adjacent to v
	Adj(v int) []int

	// ToString : return string representation
	ToString() string
}
