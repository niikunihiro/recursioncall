package main

import "fmt"

/*
目標とポイント
再帰プログラムは自分自身を繰り返し呼び出し、そのたびに異なる引数を渡していく。
再帰の仕組みと様々な再帰コードの例について学ぶ
再帰とスタックの関係を学ぶ
末尾再帰について学習する

キーワード
再帰、再帰呼出し、再帰とスタック、階乗、フィボナッチ数、GCD、フラクタル図形、末尾再帰、反復

再帰呼び出し（recursion call）とは、関数や手続きなどが自分自身を呼出し実行する事を言う
アルゴリズムの中には再帰的にコードを記述する事によって効果的な処理をできるものがある
*/

// factorial 階乗の計算
// n * (n-1) * (n-2) * ... 2 * 1
func factorial(n uint) uint {
	var i uint
	for i = 0; i < n; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("Called 'factorial(%d)' (winding) \n", n)

	if n == 0 {
		// 再帰関数の終了条件 0! = 1
		return 1
	}

	v := n * factorial(n-1)

	for i = 0; i < n; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("Returning 'factorial(%d)=%d (unwinding)' \n", n, v)

	return v
}

func main() {
	var n uint = 5
	v := factorial(n)
	fmt.Printf("Factorial of %d is %d", n, v)
}
