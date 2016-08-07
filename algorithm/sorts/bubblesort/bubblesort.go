package main

import (
	"fmt"
	"sort"
)

var data = [10]int{
	57,
	86,
	42,
	64,
	8,
	7,
	28,
	55,
	11,
	35,
}

// bubble sort
// 隣り合う要素の大小を比較しながら整列させること
// 最悪計算時間がO(n2)と遅いが、アルゴリズムが単純で実装が容易なため、
// また並列処理との親和性が高い
// 安定な内部ソート
func bubbleSort(data []int) []int {
	var tmp int
	len := len(data)

	fmt.Printf("2-before)data: %v\n", data)
	//最後の要素を除き、小->大へ
	for i := 0; i < (len - 1); i++ {
		//最後の要素-1をスタートとし、大->小へ
		for j := (len - 1); j > i; j-- {
			//スタート地点は常に同じ位置
			//ソートの終了位置が、徐々に後ろに移動していく

			//前のほうが大きい時は互いに入れ替える
			if data[j-1] > data[j] {
				tmp = data[j-1]
				data[j-1] = data[j]
				data[j] = tmp
				//golangのswapを使う
			}
		}
	}
	fmt.Printf("2-after)data: %v\n", data)
	return data
}

// main
func main() {
	//test data(int)をrandomで作成

	//とりあえず、宣言したものからsliceにコピー
	var s1 []int
	s1 = data[:]

	//sort
	fmt.Printf("1-before)s1: %v\n", s1)
	sort.Sort(sort.IntSlice(s1))
	fmt.Printf("1-after)s1: %v\n", s1)

	s2 := []int{57, 86, 42, 64, 8, 7, 28, 55, 11, 35}

	fmt.Printf("3-before)s2: %v\n", s2)
	ret := bubbleSort(s2)
	fmt.Printf("3-after)s2: %v\n", s2)
	fmt.Printf("3-after)ret: %v\n", ret)
	//sliceは参照情報なので、関数内の結果を返さなくても反映される。
}
