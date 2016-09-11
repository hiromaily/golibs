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

// selection sort
// 配列された要素から、最大値やまたは最小値を探索し配列最後の要素と入れ替えをおこなうこと
// 最悪計算時間が遅いが、単純で実装が容易。しかし、安定ソート(stable sort)ではない。
//
// データ列中で一番小さい値を探し、1番目の要素と交換する。
// 次に、2番目以降のデータ列から一番小さい値を探し、2番目の要素と交換する
func selectionSort(data []int) []int {
	var tmp, min int
	len := len(data)

	fmt.Printf("2-before)data: %v\n", data)
	//最後の要素を除く
	for i := 0; i < (len - 1); i++ {
		//minになる要素番号を精査
		min = i
		for j := i + 1; j < len; j++ {
			//iの次に前の要素から後ろまで順に比較
			//最小値を示す変数にもかわらず、min変数が大きい場合
			if data[min] > data[j] {
				min = j
			}
		}
		tmp = data[i]
		data[i] = data[min]
		data[min] = tmp
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
	ret := selectionSort(s2)
	fmt.Printf("3-after)s2: %v\n", s2)
	fmt.Printf("3-after)ret: %v\n", ret)
	//sliceは参照情報なので、関数内の結果を返さなくても反映される。
}
