package graph

/*
This undirected  Graph implementation maintains a vertex-indexed array of lists
of integers. Every edge appears twice: if an edge connects v and w, then w appears
in v’s list and v appears in w’s list.

This implemntation assumes the number of vertices are known at init time.
edges can be added dynamically

Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter four
*/

// AdjacencyListGraph : implmentation of Undirected Graph
type AdjacencyListGraph struct {
	v   int
	e   int
	adj [][]int
}

// NewAdjacencyListGraph : create and initialize
func NewAdjacencyListGraph(v int) AdjacencyListGraph {
	ag := AdjacencyListGraph{}
	ag.Graph(v)
	return ag
}

// Graph : create a graph with no edges
func (graph *AdjacencyListGraph) Graph(v int) {
	graph.v = v
	graph.e = 0
	graph.adj = make([][]int, v)
	// initialize the array of adjacency lists
	for i := range graph.adj {
		graph.adj[i] = make([]int, 0)
	}
}

// V : number of vertices
func (graph *AdjacencyListGraph) V() int {
	return graph.v
}

// E : number of edges
func (graph *AdjacencyListGraph) E() int {
	return graph.e
}

// AddEdge : add an edge from v to w
func (graph *AdjacencyListGraph) AddEdge(v int, w int) {
	graph.adj[v] = append(graph.adj[v], w)
	graph.adj[w] = append(graph.adj[w], v)
}

//Adj : get a list of vertices adjacent to v
func (graph *AdjacencyListGraph) Adj(v int) []int {
	return graph.adj[v]
}

// ToString : return string representation
func (graph *AdjacencyListGraph) ToString() string {
	return "AdjacencyListGraph"
}
