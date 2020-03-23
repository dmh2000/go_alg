package uf

import (
	"flag"
	"os"
	"testing"
)

const testDataFile string = "../../data/mediumUf.txt"

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

// -----------------------------
// Test method set of Quick-Find
// ------------------------------

func TestQf(t *testing.T) {

	// create a quick-find object
	qf := Qf{}
	qf.Uf(len(testData))

	// execute the test code, passing the
	b, p, q := testUf(&qf, testData)
	if !b {
		t.Errorf("Quick Find : should be connected but weren't p:%d q:%d", p, q)
	}
}

func BenchmarkQf(b *testing.B) {
	for i := 0; i < 1000; i++ {
		// create a quick-find object
		qf := Qf{}
		qf.Uf(len(testData))

		// assign to interface vairable so its polymorphic
		testUf(&qf, testData)
	}
}

// ------------------------------
// Test method set of Quick-Union
// -------------------------------

func TestQu(t *testing.T) {

	// create a quick-union object
	qu := Qu{}
	qu.Uf(len(testData))

	// execute the test code
	b, p, q := testUf(&qu, testData)
	if !b {
		t.Errorf("Quick Union : should be connected but weren't p:%d q:%d", p, q)
	}
}

func BenchmarkQu(b *testing.B) {
	for i := 0; i < 1000; i++ {

		// create a quick-union object
		qu := Qu{}
		qu.Uf(len(testData))

		// execute the test code
		testUf(&qu, testData)
	}
}

// -----------------------------------------
// Test method set of Weighted-Quick-Union
// -----------------------------------------

func TestWqu(t *testing.T) {
	// create a quick-union object
	wqu := Wqu{}
	wqu.Uf(len(testData))

	// execute the test code
	b, p, q := testUf(&wqu, testData)
	if !b {
		t.Errorf("Quick Find : should be connected but weren't p:%d q:%d", p, q)
	}

}

func BenchmarkWqu(b *testing.B) {
	for i := 0; i < 1000; i++ {

		// create a weighted-quick-union object
		wqu := Wqu{}
		wqu.Uf(len(testData))

		// execute the test code
		testUf(&wqu, testData)
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
