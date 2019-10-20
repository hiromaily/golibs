package fuzzysearch

import (
	lv "github.com/creasty/go-levenshtein"
)

// GetDistance is to get distance
func GetDistance(base, compared string) int {
	return lv.Distance(base, compared)
}
