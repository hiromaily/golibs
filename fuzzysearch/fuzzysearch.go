package fuzzysearch

import (
	lv "github.com/creasty/go-levenshtein"
)

func GetDistance(base, compared string) int {
	return lv.Distance(base, compared)
}
