package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var x float64
	var s string

	fmt.Printf("数を入力してください:")
	fmt.Scan(&x)

	// 絶対値を返す
	fmt.Printf("%fの絶対値は%f\n", x, math.Abs(x))
	// 切り上げ
	fmt.Printf("%.2fを切り上げた値は%.2f\n", x, math.Ceil(x))
	// 切り捨て
	fmt.Printf("%.2fを切り捨てた値は%.2f\n", x, math.Floor(x))
	// 四捨五入
	fmt.Printf("%.2fを四捨五入した値は%.2f\n", x, math.Round(x))
	// べき乗
	fmt.Printf("%.2fのeのべき乗は%.2f\n", x, math.Exp(x))
	// 平方根
	fmt.Printf("%.2fの平方根は%.2f\n", x, math.Sqrt(x))
	// 立方根
	fmt.Printf("%.2fの立方根は%.2f\n", x, math.Cbrt(x))

	fmt.Printf("文字列を入力してください:")
	fmt.Scan(&s)

	// 文字をカウント len(s)だと全角文字が含まれている場合,1文字3バイトだから正確ではない
	fmt.Printf("%sの文字数は%d\n", s, strings.Count(s, "")-1)
	// 文字列を繰り返す
	fmt.Printf("%sを2回繰り返すと%s\n", s, strings.Repeat(s, 2))
	// 全て大文字に変換
	fmt.Printf("%sの小文字を大文字にすると%s\n", s, strings.ToUpper(s))
	// 全て小文字に変換
	fmt.Printf("%sの大文字を小文字にすると%s\n", s, strings.ToLower(s))
}