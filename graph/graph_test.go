package graph

/*
	Go code in this file is based on the Java implementation from
	Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter 4
*/

import (
	"testing"
)

const tinyG string = "../data/tinyG.txt"
const mediumG string = "../data/mediumG.txt"

var testData graphData

func TestDfsTiny(t *testing.T) {
	var g Graph
	f := tinyG
	// get the test data
	testData = getTestData(f)

	// create an empty graph
	g = NewAdjacencyListGraph(testData.v)
	// load the data
	for i := range testData.edges {
		v := testData.edges[i].v
		w := testData.edges[i].w
		g.AddEdge(v, w)
	}
	// perform the DFS
	source := 0
	d := NewDFS(g, source)
	// check what should be marked
	for i := 0; i < 7; i++ {
		if !d.HasPathTo(i) {
			t.Errorf("DFS %s:%d node %d should be marked", f, source, i)
		}
	}
	// check what should not be marked
	for i := 7; i < g.V(); i++ {
		if d.HasPathTo(i) {
			t.Errorf("DFS %s:%d node %d should not be marked", f, source, i)
		}
	}
	// print the path
	path := d.PathTo(1)
	if len(path) != 2 {
		t.Error("DFS path is wrong length")
	}
	path = d.PathTo(2)
	if len(path) != 2 {
		t.Error("DFS path is wrong length")
	}
	path = d.PathTo(3)
	if len(path) != 4 {
		t.Error("DFS path is wrong length")
	}
	path = d.PathTo(4)
	if len(path) != 3 {
		t.Error("DFS path is wrong length")
	}
	path = d.PathTo(5)
	if len(path) != 2 {
		t.Error("DFS path is wrong length")
	}
	path = d.PathTo(6)
	if len(path) != 4 {
		t.Error("DFS path is wrong length")
	}
}

func TestDfsMedium(t *testing.T) {
	f := mediumG
	// get the test data
	testData = getTestData(f)

	// create an empty graph
	g := NewAdjacencyListGraph(testData.v)
	// load the data
	for i := range testData.edges {
		v := testData.edges[i].v
		w := testData.edges[i].w
		g.AddEdge(v, w)
	}
	// perform the DFS
	source := 0
	d := NewDFS(g, source)

	// the medium dataset is fully connected so there
	// should be paths to all vertices
	for w := 1; w < g.V(); w++ {
		if !d.HasPathTo(w) {
			t.Errorf("Medium has not path from %d to %d", source, w)
		}
	}
}

func TestBfsTiny(t *testing.T) {
	f := tinyG

	// get the test data
	testData = getTestData(f)

	// create an empty graph
	g := NewAdjacencyListGraph(testData.v)
	// load the data
	for i := range testData.edges {
		v := testData.edges[i].v
		w := testData.edges[i].w
		g.AddEdge(v, w)
	}
	// perform the DFS
	source := 0
	bfs := NewBFS(g, source)
	// check what should be marked
	for i := 0; i < 7; i++ {
		if !bfs.HasPathTo(i) {
			t.Errorf("BFS %s:%d node %d should be marked", f, source, i)
		}
	}
	// check what should not be marked
	for i := 7; i < g.V(); i++ {
		if bfs.HasPathTo(i) {
			t.Errorf("BFS %s:%d node %d should not be marked", f, source, i)
		}
	}
	// print the path
	path := bfs.PathTo(1)
	if len(path) != 2 {
		t.Error("BFS path is wrong length")
	}
	path = bfs.PathTo(2)
	if len(path) != 2 {
		t.Error("BFS path is wrong length")
	}
	path = bfs.PathTo(3)
	if len(path) != 3 {
		t.Error("BFS path is wrong length")
	}
	path = bfs.PathTo(3)
	if len(path) != 3 {
		t.Error("BFS path is wrong length")
	}
	path = bfs.PathTo(2)
	if len(path) != 2 {
		t.Error("BFS path is wrong length")
	}
	path = bfs.PathTo(2)
	if len(path) != 2 {
		t.Error("BFS path is wrong length")
	}
}
