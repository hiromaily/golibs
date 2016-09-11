package bulkdata

import (
	"math/rand"
	"time"
)

// MakeIntData is to return array from random number
func MakeIntData(num int) (values []int) {
	// UNIX 時間をシードにして乱数生成器を用意する
	t := time.Now().Unix()
	s := rand.NewSource(t)
	r := rand.New(s)

	// ランダムな値の入った配列を作る
	values = make([]int, num)
	for i := 0; i < num; i++ {
		values[i] = r.Intn(num)
	}

	return
}
