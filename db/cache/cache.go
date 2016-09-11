package cache

import (
	"fmt"
)

//cahce
var cacheData map[string][]map[string]interface{}

// Get is to get data by key
func Get(key string) ([]map[string]interface{}, error) {
	if value, ok := cacheData[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("data didn't find on cache, key is %s", key)
}

// Set is to set data by key
func Set(key string, data []map[string]interface{}) {
	cacheData[key] = data
}
