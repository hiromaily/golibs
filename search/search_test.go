package search_test

import (
	"testing"

	. "github.com/hiromaily/golibs/search"
)

func createTestData() []int {
	maxNum := 10000
	list := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		list[i] = i + 1
	}
	return list
}

// set even number
func createTestData2() []int {
	maxNum := 10000
	list := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		list[i] = 100 + (i * 2)
	}
	return list
}

func TestSearch(t *testing.T) {
	t.SkipNow()
	list := createTestData()
	testTargets := []int{
		157,
		3978,
		5002,
		7888,
		9923,
	}
	//Search() took 771ns
	//BinarySearch() took 361ns

	//Search() took 3.669µs
	//BinarySearch() took 393ns

	//Search() took 4.374µs
	//BinarySearch() took 290ns

	//Search() took 7.033µs
	//BinarySearch() took 290ns

	//Search() took 8.565µs
	//BinarySearch() took 205ns

	for _, target := range testTargets {
		if !Search(target, list) {
			t.Errorf("target %d is not found", target)
		}
		if !BinarySearch(target, list) {
			t.Errorf("target %d is not found", target)
		}
	}
}

func TestSearchNearest(t *testing.T) {
	list := createTestData2()
	//should set odd number
	testTargets := []int{
		4,
		313,
		7981,
		10019,
		15888,
		19925,
	}

	//SearchNearest() took 396ns
	//BinarySearchNearest() took 320ns

	//SearchNearest() took 2.524µs
	//BinarySearchNearest() took 229ns

	//SearchNearest() took 3.149µs
	//BinarySearchNearest() took 257ns

	//SearchNearest() took 4.82µs
	//BinarySearchNearest() took 256ns

	//SearchNearest() took 6.006µs
	//BinarySearchNearest() took 171ns

	for _, target := range testTargets {
		val, idx := SearchNearest(target, list)
		t.Logf("nearest value by SearchNearest(): %d, index: %d", val, idx)
		val2, idx2 := BinarySearchNearest(target, list)
		t.Logf("nearest value by BinarySearchNearest(): %d, index: %d", val2, idx2)
	}
}
