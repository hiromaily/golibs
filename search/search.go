package search

import (
	"time"

	tm "github.com/hiromaily/golibs/time"
)

func Search(target int, list []int) bool {
	defer tm.Track(time.Now(), "Search()")

	for _, val := range list {
		if target == val {
			return true
		}
	}
	return false
}

// get nearest and smaller one
func SearchNearest(target int, list []int) (int, int) {
	defer tm.Track(time.Now(), "SearchNearest()")

	if list[0] > target {
		return 0, 0
	}

	for idx, val := range list {
		if target < val {
			return list[idx-1], idx - 1
		}
	}
	return 0, 0
}

func BinarySearch(target int, list []int) bool {
	defer tm.Track(time.Now(), "BinarySearch()")

	low := 0
	high := len(list) - 1

	for low <= high {
		median := (low + high) / 2

		if list[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(list) || list[low] != target {
		return false
	}

	return true
}

// get nearest and smaller one
func BinarySearchNearest(target int, list []int) (int, int) {
	defer tm.Track(time.Now(), "BinarySearchNearest()")

	if list[0] > target {
		return 0, 0
	}

	low := 0
	high := len(list) - 1

	for low <= high {
		median := (low + high) / 2

		if list[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	//fmt.Println(low, high, len(list), list[low], target)
	//157 156 10000 314 313
	if low == len(list) || list[low] != target {
		return list[low-1], low - 1
	}

	return list[low], low
}
