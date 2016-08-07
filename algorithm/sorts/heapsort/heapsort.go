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

// heap sort(未実装)
// リストの並べ替えを二分ヒープ木を用いて行うソートのアルゴリズム
// 安定ソート(stable sort)ではない。
//
// 1.未整列のリストから要素を取り出し、順にヒープに追加する。すべての要素を追加するまで繰り返し。
// 2.ルート（最大値または最小値)を取り出し、整列済みリストに追加する。すべての要素を取り出すまで繰り返し
// http://www.th.cs.meiji.ac.jp/researches/2005/omoto/heapsort.html
func heapSort(data []int) []int {
	//var tmp, min int
	//len := len(data)

	fmt.Printf("2-before)data: %v\n", data)

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
	ret := heapSort(s2)
	fmt.Printf("3-after)s2: %v\n", s2)
	fmt.Printf("3-after)ret: %v\n", ret)
	//sliceは参照情報なので、関数内の結果を返さなくても反映される。
}
