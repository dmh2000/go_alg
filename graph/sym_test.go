package graph

/*
	Go code in this file is based on the Java implementation from
	Sedgewick, Robert. Algorithms . Pearson Education. Kindle Edition. Chapter 4
*/

import (
	"fmt"
	"testing"
)

const routes string = "../data/routes.txt"
const movies string = "../data/movies.txt"

func TestSymbolGraph1(t *testing.T) {
	sg := NewSymbolGraph(routes, " ")

	g := sg.G()

	k := sg.keys

	for _, s := range k {
		fmt.Println(s)
		fmt.Printf("  ")
		for _, v := range g.Adj(sg.Index(s)) {
			fmt.Printf("%v:%v, ", v, sg.Name(v))
		}
		fmt.Println()
	}
}

var test_name string

func TestSymbolGraph2(t *testing.T) {
	sg := NewSymbolGraph(movies, "/")

	g := sg.G()

	k := sg.keys

	// loop through all keys to make sure it doesn't crash
	for _, s := range k {
		for _, v := range g.Adj(sg.Index(s)) {
			test_name = sg.Name(v)
		}
	}

	// print an instance of a p ath
	s := "'Breaker' Morant (1980)"
	fmt.Println(s)
	for _, v := range g.Adj(sg.Index(s)) {
		fmt.Println("  " + sg.Name(v))
	}
}

func TestDofS(t *testing.T) {
	DegreesOfSeparation(
		"Bacon, Kevin",
		"Russell, Kurt (I)",
		movies,
		"/",
	)

}
