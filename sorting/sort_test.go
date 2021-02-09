package algsort

import (
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

func (data dataSlice) get(index int) int {
	return data[index]
}

func (data dataSlice) set(index int, value int) {
	data[index] = value
}

func (data dataSlice) clone() AlgSort {
	return make(dataSlice,len(data))
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

func doSort(f func(AlgSort),data dataSlice, t *testing.T) {
	f(data)

	if !data.isSorted() {
		t.Errorf("data is not sorted")
	}
}

func randomData(k int) dataSlice {
	data := make(dataSlice, k)

	for j := 0; j < data.Len(); j++ {
		data.set(j, rand.Int())
	}

	return data
}

func descendingData(k int) dataSlice {
	data := make(dataSlice, k)

	j := data.Len() - 1
	for i := range data {
		data[i] = j
		j--
	}

	return data
}

func ascendingData(k int) dataSlice {
	data := make(dataSlice, k)

	for j := range data {
		data[j] = j
	}
	return data
}

var iterations = 1
var dataLength = 1024 * 1024

func TestMergeSortAscending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(MergeSort, ascendingData(dataLength), t)
	}
}

func TestMergeSortDescending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(MergeSort, descendingData(dataLength), t)
	}
}

func TestMergeSortRandom(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(MergeSort,randomData(dataLength), t)
	}
}

func TestQuickSortAscending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(QuickSort, ascendingData(dataLength), t)
	}
}

func TestQuickSortDescending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(QuickSort,descendingData(dataLength), t)
	}
}

func TestQuickSortRandom(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(QuickSort, randomData(dataLength), t)
	}
}

func TestGoSortAscending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(GoSort, ascendingData(dataLength), t)
	}
}

func TestGoSortDescending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(GoSort, descendingData(dataLength), t)
	}
}

func TestGoSortRandom(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(GoSort, randomData(dataLength), t)
	}
}

func TestInsertionSortAscending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(InsertionSort, ascendingData(dataLength), t)
	}
}

func TestInsertionSortDescending(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(InsertionSort, descendingData(dataLength), t)
	}
}

func TestInsertionSortRandom(t *testing.T) {
	for i:=0;i<iterations;i++ {
		doSort(InsertionSort, randomData(dataLength), t)
	}
}

func TestMain(m *testing.M) {
	// dataLength = 1024 * 1024
	// iterations = 5
	os.Exit(m.Run())
}
