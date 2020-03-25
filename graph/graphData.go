package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Edge : data for one edge
type Edge struct {
	v int
	w int
}

// graphData : test data input
type graphData struct {
	v     int
	e     int
	edges []Edge
}

func scanInt(fs *bufio.Scanner) int {
	// scan next token
	b := fs.Scan()
	if !b {
		panic("end of file too soon")
	}
	// get the string
	s := fs.Text()
	// parse it
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err.Error() + ":" + s)
	}

	return int(i)
}

func scanEdge(fs *bufio.Scanner) (int, int) {
	var v int
	var w int

	// scan next token
	b := fs.Scan()
	if !b {
		panic("end of file too soon")
	}
	// get the string
	s := fs.Text()
	// parse it
	n, err := fmt.Sscanf(s, "%d %d", &v, &w)
	if err != nil {
		panic(err.Error() + ":" + s)
	}
	if n != 2 {
		panic("less than 2 tokens :" + s)
	}

	return v, w
}

// GetTestData get test data for union connectivity algorithms
func getTestData(fname string) graphData {
	var vertices int
	var edges int

	f, err := os.Open(fname)
	if err != nil {
		panic("can't open file " + fname)
	}
	defer f.Close()
	fs := bufio.NewScanner(f)

	// get number of vertices
	vertices = scanInt(fs)
	if vertices <= 0 {
		panic("invalid vertices value : " + string(vertices))
	}

	// get number of edges
	edges = scanInt(fs)
	if edges <= 0 {
		panic("invalid edges value : " + string(edges))
	}

	// create a graph data object
	data := graphData{
		v:     vertices,
		e:     edges,
		edges: make([]Edge, edges),
	}

	// read all the edges
	for i := 0; i < edges; i++ {
		v, w := scanEdge(fs)
		data.edges[i] = Edge{v, w}
	}

	return data
}
