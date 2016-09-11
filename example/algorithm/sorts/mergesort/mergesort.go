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

// merge sort
// 既に整列してある複数個の列を1個の列にマージする際に、小さいものから先に新しい列に並べれば、
// 新しい列も整列されている、というボトムアップの分割統治法によるもの。
// 大きい列を多数の列に分割し、そのそれぞれをマージする作業は並列化できる。
// 安定な内部ソート
// （ナイーブな）クイックソートと比べると、最悪計算量は少ない。
// ランダムなデータでは通常、クイックソートのほうが速い。
// 1.データ列を分割する（通常、等分する）
// 2.各々をソートする (マージソートを再帰的に適用する)
// 3.二つのソートされたデータ列をマージする
func mergeSort(values []int) (ret []int) {
	fmt.Printf("2-before)data: %v\n", values)

	//基本的に分割と、sortを再帰的に実行することで実現する。
	left, right := split(values)
	ret = sortint(left, right)

	fmt.Printf("2-after)data: %v\n", ret)
	return
}

// 1.split
func split(values []int) (left, right []int) {
	// スライスを真ん中でふたつに分割する
	left = values[:len(values)/2]
	right = values[len(values)/2:]
	return
}

// 2.sort
func sortint(left, right []int) (ret []int) {
	//2つのスライスをそれぞれ再帰的にソートする
	//(sortintメソッドないで、
	if len(left) > 1 {
		l, r := split(left)
		left = sortint(l, r)
	}
	//2つのスライスをそれぞれ再帰的にソートする
	if len(right) > 1 {
		l, r := split(right)
		right = sortint(l, r)
	}

	// ソート済みのふたつのスライスをひとつにマージする
	ret = merge(left, right)
	return
}

// 3.merge
func merge(left, right []int) (ret []int) {
	ret = []int{}
	for len(left) > 0 && len(right) > 0 {
		var x int
		// ソート済みの2つのスライスからより小さいものを選んで追加していく
		if right[0] > left[0] {
			x, left = left[0], left[1:]
		} else {
			x, right = right[0], right[1:]
		}
		ret = append(ret, x)
	}
	ret = append(ret, left...)
	ret = append(ret, right...)
	return
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
	ret := mergeSort(s2)
	fmt.Printf("3-after)s2: %v\n", s2)
	fmt.Printf("3-after)ret: %v\n", ret)
	//sliceは参照情報なので、関数内の結果を返さなくても反映される。
}
