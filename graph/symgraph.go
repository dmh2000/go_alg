package graph

/*
	Go code in this file is based on the Java implementation from
	Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter 4
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// public class SymbolGraph
// {
//    private ST<String, Integer> st;              // String -> index
//    private String[] keys;                       // index -> String
//    private Graph G;                             // the graph

// SymbolGraph : graph with information attached to nodes
type SymbolGraph struct {
	st   map[string]int
	keys []string
	g    Graph
}

// NewSymbolGraph : creates an empty SymbolGraph
func NewSymbolGraph(stream string, sp string) SymbolGraph {
	sg := SymbolGraph{}
	sg.st = make(map[string]int)

	// build index by reading strings from input file
	f, err := os.Open(stream)
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(f)
	index := 0
	// read one line at a time and split
	// associate new strings with an index
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, sp)
		for _, word := range words {
			// is word already in the map?
			_, ok := sg.st[word]
			if !ok {
				sg.st[word] = index
				index++
			}
		}
	}

	// create inverted index
	sg.keys = make([]string, index)
	// loop over keys in the symbol table
	for k := range sg.st {
		// add a key with its index to the array
		sg.keys[sg.st[k]] = k
	}

	// create a graph
	sg.g = NewAdjacencyListGraph(len(sg.keys))

	// second pass
	// build index by reading strings from input file
	f, err = os.Open(stream)
	if err != nil {
		panic(err.Error())
	}

	scanner = bufio.NewScanner(f)
	index = 0
	// read one line at a time and split
	// associate new strings with an index
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, sp)
		v := sg.st[words[0]]
		for i := 1; i < len(words); i++ {
			sg.g.AddEdge(v, sg.st[words[i]])
		}
	}

	return sg
}

// Contains : graph has the string value
func (sg *SymbolGraph) Contains(s string) bool {
	_, ok := sg.st[s]
	return ok
}

// Index : returns the index of the string in the graph
func (sg *SymbolGraph) Index(s string) int {
	i, ok := sg.st[s]
	if !ok {
		panic("string s not in symbol table")
	}
	return i
}

// Name : returns the string associated with the index i
func (sg *SymbolGraph) Name(v int) string {
	if v >= len(sg.keys) {
		panic("Index out of bounds for keys")
	}
	return sg.keys[v]
}

// G : returns the underlying graph
func (sg *SymbolGraph) G() Graph {
	return sg.g
}

// ================================================
// DEGREES OF SEPARATION
// ================================================

// DegreesOfSeparation : find separation of source to sink
func DegreesOfSeparation(source string, sink string, fname string, sep string) {
	sg := NewSymbolGraph(fname, sep)
	g := sg.G()

	if !sg.Contains(source) {
		fmt.Printf("%s no in database\n", source)
		return
	}

	s := sg.Index(source)

	bfs := NewBFS(g, s)

	if sg.Contains(sink) {
		t := sg.Index(sink)
		if bfs.HasPathTo(t) {
			indent := ""
			for _, v := range bfs.PathTo(t) {
				fmt.Println(indent + sg.Name(v))
				indent += "  "
			}
		} else {
			fmt.Println("not connected")
		}

	} else {
		fmt.Printf("%s not in database\n", sink)
	}
}
