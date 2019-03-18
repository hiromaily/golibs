package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func syncMap(smap *sync.Map) {
	//key value
	smap.Store("key1", "aaaa")
	smap.Store("key2", "bbbb")
	smap.Store(1, 2)

	smap.LoadOrStore(2, 4)

	//delete
	smap.Delete("key2")

	//get
	val, ok := smap.Load("key1")
	if ok {
		fmt.Printf("key: %s, value: %v\n", "key1", val)
	}

	smap.Range(func(key, val interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, val)
		return true
	})
}

func SyncMap() {
	smap := new(sync.Map)

	for i := 0; i < 5; i++ {
		go syncMap(smap)
	}
}

var id int64 = 1

func incrementID() {
	atomic.AddInt64(&id, 1)
}

func SyncAtomic() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			incrementID()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("id: %d\n", id)
}
