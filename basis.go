package main

import (
	"fmt"
)

func main() {
	// 宣言した後、値を代入パターン
var msg string
msg = "hello world"

// 宣言と代入を一緒にするパターン

var msg string = "Hello World"

// 宣言と代入を一緒にするパターン (型省略可能)
var msg = "Hello Hello"

// 宣言と代入を一緒にするパターン (var省略)
msg := "Super Hello"

// 複数の同じ型の変数を同時に定義
var a, b int
a, b = 10, 15

// 複数の同じ型の変数を同時に定義 (var省略)
a, b := 10, 14

// 複数の型違いの変数を同時に定義
var (
  c int
  d string
)

c = 20
d = "hoge"

// 複数の型違いの変数を同時に定義 (型省略パターン)

var (
  c = 20
  d = "hoge"
)

//定数の定義
const (
	sun = iota // 0
	mon // 1
	tue // 2
  )

  // 引数なし関数
func sayHi() {
    fmt.Println("hi!")
}

// 引数あり関数
func sayName(name string) {
    fmt.Println(name)
}

// 名前付きreturn関数
// 関数内で使った変数名を返す
func getHelloMessage(name string) (msg string) {
    msg = "Hello " + name
    return
}

// 変数に関数を格納
sayHi()
sayName("gcfuji")
fmt.Println(getMessage("gcfuji"))
fmt.Println(getHelloMessage("Gemcook"))

// 複数の返り値を返す事ができる
func swap(a int, b int) (int, int) {
    return b, a
}

fmt.Println(swap(5, 2))

// 関数もデータ型なので、変数に代入可能
// その際は関数名はいらない
f := func(a int, b int) (int, int) {
    return b, a
}
fmt.Println(f(3, 8))

// 即時関数的な事も可能
func(msg string) {
    fmt.Println(msg)
("Fujimoto")

// 宣言と代入を分ける
var a [5]int // a[0] - a[4]

a[2] = 3
a[4] = 10
fmt.Println(a)

// 宣言と代入を同時にする
b := [3]int{1, 3, 6}
fmt.Println(b)

// 配列の個数が未定の場合
c := [...]int{2, 4, 7, 5, 5}
fmt.Println(c)

}