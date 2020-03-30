package algsort

import (
	"flag"
	"math/rand"
	"os"
	"testing"
)

type dataSlice []int

func (data dataSlice) Less(a int, b int) bool {
	return data[a] < data[b]
}

func (data dataSlice) Swap(a int, b int) {
	data[a], data[b] = data[b], data[a]
}

func (data dataSlice) Len() int {
	return len(data)
}

func (data dataSlice) lt(a, b int) bool {
	return data[a] < data[b]
}

func (data dataSlice) exch(a, b int) {
	data[a], data[b] = data[b], data[a]
}

func (data dataSlice) length() int {
	return len(data)
}

func (data dataSlice) isSorted() bool {
	j := 0
	for i := 1; i < len(data); i++ {
		if data[j] > data[i] {
			return false
		}
		j = i
	}
	return true
}

func TestInsertionSort(t *testing.T) {
	var d dataSlice

	d = make([]int, 256)

	for i := range d {
		d[i] = rand.Int()
	}

	InsertionSort(d)

	prev := 0
	for i := 1; i < len(d); i++ {
		if d[prev] > d[i] {
			t.Errorf("%v:%v should be less than %v:%v", prev, d[prev], i, d[i])
		}
		prev = i
	}
}

func TestQuickSort(t *testing.T) {
	var d dataSlice

	d = make([]int, 256)

	for i := range d {
		d[i] = rand.Int()
	}

	QuickSort(d)

	prev := 0
	for i := 1; i < len(d); i++ {
		if d[prev] > d[i] {
			t.Errorf("%v:%v should be less than %v:%v", prev, d[prev], i, d[i])
			return
		}
		prev = i
	}
}

func TestMain(m *testing.M) {
	// activate benchmarking if required
	testing.Init()

	// parse command line flags for test
	flag.Parse()

	// run and exit
	os.Exit(m.Run())
}
