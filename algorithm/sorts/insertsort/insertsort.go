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

// insert sort
// 整列してある配列に追加要素を適切な場所に挿入すること
// 平均計算時間・最悪計算時間がともにO(n2)と遅いが、アルゴリズムが単純で実装が容易
// 安定な内部ソート
func insertSort(data []int) []int {
	var tmp, j int
	len := len(data)

	fmt.Printf("2-before)data: %v\n", data)
	for i := 1; i < len; i++ {
		tmp = data[i]
		//隣接する前後の番号を比較
		if data[i-1] > tmp {
			//indexの小さい側の集合体と、indexの大きい側の集合体に分けられる。
			//ここでの比較はその隣接点にあたる
			//indexの小さい側のほうがデータが大きい場合、処理が発生
			j = i
			for {
				//ここからは、indexの小さい側の集合体に対しての処理(こちらの集合はソート済み)
				//既にdata[i]は退避済みなので、indexの小さい方を大きい方にセットする
				data[j] = data[j-1]
				j--
				//Jが0以上、かつ、さらに下のindexに対して、退避データを比較して、大きければ処理を継続
				if j > 0 && data[j-1] > tmp {
					continue
				} else {
					break
				}
			}
			//小さい集合内における、適切な位置へのソートが完了したので、差し込む。
			data[j] = tmp
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
	ret := insertSort(s2)
	fmt.Printf("3-after)s2: %v\n", s2)
	fmt.Printf("3-after)ret: %v\n", ret)
	//sliceは参照情報なので、関数内の結果を返さなくても反映される。
}
