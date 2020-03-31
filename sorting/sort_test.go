package algsort

import (
	"math/rand"
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

func (data dataSlice) get(index int) int {
	return data[index]
}

func (data dataSlice) set(index int, value int) {
	data[index] = value
}

func (data dataSlice) clone() AlgSort {
	a := make([]int, data.length())
	return dataSlice(a)
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

const iterations = 4096

func Random(f func(AlgSort), t *testing.T) {
	var data dataSlice

	for k := 1; k < iterations; k++ {
		data = make([]int, k)

		for j := 0; j < data.length(); j++ {
			data.set(j, rand.Int())
		}

		f(data)

		if !data.isSorted() {
			t.Errorf("%v : data is not sorted", k)
		}
	}
}

func Inverted(f func(AlgSort), t *testing.T) {
	var data dataSlice

	for k := 1; k < iterations; k++ {
		data = make([]int, k)

		j := data.length() - 1
		for i := range data {
			data[i] = j
			j--
		}

		f(data)

		if !data.isSorted() {
			t.Error("data is not sorted")
		}

	}
}

func Inorder(f func(AlgSort), t *testing.T) {
	var data dataSlice
	for k := 1; k < iterations; k++ {
		data = make([]int, k)

		for j := range data {
			data[j] = j
		}

		f(data)

		if !data.isSorted() {
			t.Error("data is not sorted")
		}
	}
}

func TestMergeSortInorder(t *testing.T) {
	Inorder(MergeSort, t)
}

func TestMergeSortInverted(t *testing.T) {
	Inverted(MergeSort, t)
}

func TestMergeSortRandom(t *testing.T) {
	Random(MergeSort, t)
}

func TestQuickSortInorder(t *testing.T) {
	Inorder(QuickSort, t)
}

func TestQuickSortInverted(t *testing.T) {
	Inverted(QuickSort, t)
}

func TestQuickSortRandom(t *testing.T) {
	Random(QuickSort, t)
}

func TestInsertionSortInorder(t *testing.T) {
	Inorder(InsertionSort, t)
}

func TestInsertionSortInverted(t *testing.T) {
	Inverted(InsertionSort, t)
}

func TestInsertionSortRandom(t *testing.T) {
	Random(InsertionSort, t)
}
