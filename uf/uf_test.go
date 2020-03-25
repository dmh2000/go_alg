package uf

import (
	"flag"
	"os"
	"testing"
)

const testDataFile string = "../data/mediumUf.txt"

var testData []Pair

// testUF : test an implementation of UF union-find
// input parameter uf can be any method set that implements
// the UF interface
func testUf(uf UF, data []Pair) (bool, int, int) {
	// add all the data
	for i := range data {
		p := data[i].P
		q := data[i].Q
		if uf.Connected(p, q) {
			// already connected, skip rest
			continue
		}
		uf.Union(p, q)
	}

	// now test the same data
	// all individual entries should be connected
	for i := range data {
		p := data[i].P
		q := data[i].Q
		if !uf.Connected(p, q) {
			return false, p, q
		}
	}
	return true, 0, 0
}

// testUF2 : test an implementation of UF union-find
// input parameter uf can be any method set that implements
// the UF interface
func testUf2(uf1 UF, uf2 UF, data []Pair) (bool, int, int, bool, bool) {
	// add all the data
	for i := range data {
		p := data[i].P
		q := data[i].Q
		if !uf1.Connected(p, q) {
			// not connected, add
			uf1.Union(p, q)
		}
		if !uf2.Connected(p, q) {
			// not connected, add
			uf2.Union(p, q)
		}
	}

	// now test connections in both from 0 .. length(data)
	for i := 0; i < len(data); i++ {
		// get p from first column
		p := data[i].P
		for j := 0; j < len(data); j++ {
			// get q from second column
			q := data[j].Q

			// get connected status from both algorithms
			b1 := uf1.Connected(p, q)
			b2 := uf2.Connected(p, q)

			// if they don't match its a fail
			if b1 != b2 {
				return false, p, q, b1, b2
			}
		}
	}
	return true, 0, 0, true, true
}

// -----------------------------
// Test method set of Quick-Find
// ------------------------------

func TestQf(t *testing.T) {

	// create a quick-find object
	qf := NewQuickFind(len(testData))

	// execute the test code, passing the
	b, p, q := testUf(&qf, testData)
	if !b {
		t.Errorf("Quick Find : should be connected but weren't p:%d q:%d", p, q)
	}
}

func BenchmarkQf(b *testing.B) {
	for i := 0; i < 1000; i++ {
		// create a quick-find object
		qf := NewQuickFind(len(testData))

		// assign to interface vairable so its polymorphic
		testUf(&qf, testData)
	}
}

// ------------------------------
// Test method set of Quick-Union
// -------------------------------

func TestQu(t *testing.T) {

	// create a quick-union object
	qu := NewQuickUnion(len(testData))

	// execute the test code
	b, p, q := testUf(&qu, testData)
	if !b {
		t.Errorf("Quick Union : should be connected but weren't p:%d q:%d", p, q)
	}
}

func BenchmarkQu(b *testing.B) {
	for i := 0; i < 1000; i++ {

		// create a quick-union object
		qu := NewQuickUnion(len(testData))

		// execute the test code
		testUf(&qu, testData)
	}
}

// -----------------------------------------
// Test method set of Weighted-Quick-Union
// -----------------------------------------

func TestWqu(t *testing.T) {
	// create a weighted-quick-union object
	wqu := NewWeightedQuickUnion(len(testData))

	// execute the test code
	b, p, q := testUf(&wqu, testData)
	if !b {
		t.Errorf("Quick Find : should be connected but weren't p:%d q:%d", p, q)
	}

}

func BenchmarkWqu(b *testing.B) {
	for i := 0; i < 1000; i++ {

		// create a weighted-quick-union object
		wqu := NewWeightedQuickUnion(len(testData))

		// execute the test code
		testUf(&wqu, testData)
	}
}

// TestQfWqu : check that quick-find and weighted-quick-union
// results match for all combinations of data
func TestQfWqu(t *testing.T) {
	// create a  quick-find object
	qf := NewQuickFind(len(testData))

	// create a quick-union object
	wqu := NewWeightedQuickUnion(len(testData))

	// execute the test code
	b, p, q, b1, b2 := testUf2(&qf, &wqu, testData)
	if !b {
		// there is a mismatch
		t.Errorf("QF vs WGU : result mismatch p:%d q:%d b1:%v b2:%v", p, q, b1, b2)
	}
}

func TestMain(m *testing.M) {
	// activate benchmarking if required
	testing.Init()

	// parse command line flags for test
	flag.Parse()

	// get the test data
	testData = GetTestData(testDataFile)

	os.Exit(m.Run())
}
