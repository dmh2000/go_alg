package main

import (
	"fmt"
	"os"
)

type xyz struct {
	s string
}

func a(x xyz, i int) {
	fmt.Println(x, i)
}

func b(x *xyz, i int) {
	a(*x, i)
}

func main() {
	// x := xyz{s: os.Args[1]}
	// i := 1
	// b(&x, i)

	b(&xyz{s: os.Args[1]}, 1)
}
