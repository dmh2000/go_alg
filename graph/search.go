package graph

/*
	Go code in this file is based on the Java implementation from
	Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter 4
*/

import (
	"container/list"
)

// DFS : depth first search on an undirected graph
type DFS struct {
	marked []bool
	edgeTo []int
	source int
}

// NewDFS : create a DFS object and perform a search on it
func NewDFS(g Graph, source int) DFS {
	// create the object
	d := DFS{}
	d.marked = make([]bool, g.V())
	d.edgeTo = make([]int, g.V())
	d.source = source

	// execute the search
	d.search(g, source)

	// return the completed object
	return d
}

// search : private function that
func (dfs *DFS) search(g Graph, v int) {
	dfs.marked[v] = true
	for _, w := range g.Adj(v) {
		if !dfs.marked[w] {
			dfs.edgeTo[w] = v
			dfs.search(g, w)
		}
	}

}

// HasPathTo : return true if there is a path from source to v
func (dfs *DFS) HasPathTo(v int) bool {
	return dfs.marked[v]
}

// PathTo : return the path followed from source to v
func (dfs *DFS) PathTo(v int) []int {
	if !dfs.HasPathTo(v) {
		return nil
	}

	path := make([]int, 0)
	for x := v; x != dfs.source; x = dfs.edgeTo[x] {
		path = append([]int{x}, path...)
	}
	path = append([]int{dfs.source}, path...)

	return path
}

// Queue : queue object
type Queue struct {
	queue list.List
}

func (q *Queue) enqueue(v int) {
	q.queue.PushBack(v)
}

func (q *Queue) dequeue() int {
	// get front element
	p := q.queue.Front()
	// get its value with type assertion
	// note : will panic if p.Value is not an int
	v := p.Value.(int)
	// remove the element
	q.queue.Remove(p)
	// return the value
	return v
}

func (q *Queue) empty() bool {
	return q.queue.Len() == 0
}

func newQueue() Queue {
	queue := Queue{queue: list.List{}}
	return queue
}

// BFS : object data for breadth-first-search algorithm
type BFS struct {
	marked []bool
	edgeTo []int
	source int
}

// NewBFS : create a DFS object and perform a search on it
func NewBFS(g Graph, source int) BFS {
	// create the object
	bfs := BFS{}
	bfs.marked = make([]bool, g.V())
	bfs.edgeTo = make([]int, g.V())
	bfs.source = source
	// execute the search
	bfs.search(g, source)
	// return the completed object
	return bfs
}

func (bfs *BFS) search(g Graph, s int) {
	queue := newQueue()
	bfs.marked[s] = true // mark the source
	queue.enqueue(s)
	for !queue.empty() {
		v := queue.dequeue()
		for _, w := range g.Adj(v) {
			if !bfs.marked[w] {
				bfs.edgeTo[w] = v
				bfs.marked[w] = true
				queue.enqueue(w)
			}
		}
	}
}

// HasPathTo : return true if there is a path from source to v
func (bfs *BFS) HasPathTo(v int) bool {
	return bfs.marked[v]
}

// PathTo : return the path from source to v
func (bfs *BFS) PathTo(v int) []int {
	if !bfs.HasPathTo(v) {
		return nil
	}

	path := make([]int, 0)
	for x := v; x != bfs.source; x = bfs.edgeTo[x] {
		// stack : push onto front
		path = append([]int{x}, path...)
	}
	// stack : push onto front
	path = append([]int{bfs.source}, path...)

	return path
}
