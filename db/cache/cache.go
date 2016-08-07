package cache

import (
	"fmt"
)

//cahce
var cacheData map[string][]map[string]interface{}

func Get(key string) ([]map[string]interface{}, error) {
	if value, ok := cacheData[key]; ok {
		return value, nil
	} else {
		return nil, fmt.Errorf("data didn't find on cache, key is %s", key)
	}
}

func Set(key string, data []map[string]interface{}) {
	cacheData[key] = data
}
