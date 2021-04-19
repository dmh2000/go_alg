package uf

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func scanPair(fs *bufio.Scanner) (int, int) {
	var p int
	var q int

	// scan next token
	b := fs.Scan()
	if !b {
		panic("end of file too soon")
	}
	// get the string
	s := fs.Text()
	// parse it
	n, err := fmt.Sscanf(s, "%d %d", &p, &q)
	if err != nil {
		panic(err.Error() + ":" + s)
	}
	if n != 2 {
		panic("less than 2 tokens :" + s)
	}

	return p, q
}

// GetTestData get test data for union connectivity algorithms
func GetTestData(fname string) []Pair {
	var count int
	var p int
	var q int

	f, err := os.Open(fname)
	if err != nil {
		panic("can't open file " + fname)
	}
	defer f.Close()
	fs := bufio.NewScanner(f)

	// get the count of elements
	count = scanInt(fs)
	if count <= 0 {
		panic("invalid count value : " + fmt.Sprintf("%v", count))
	}

	data := make([]Pair, count)
	for i := 0; i < count; i++ {
		p, q = scanPair(fs)
		data[i] = Pair{p, q}
	}

	return data
}
