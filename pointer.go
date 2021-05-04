// hello.go
package main

import (
	"fmt"
)

func main() {
	var a int
	var p *int
	a = 100
	p = &a // "&"はアドレス演算子 "a"のアドレスを調べて,その値を"p"に代入
	fmt.Printf(*p) // "*"はアドレスの中身 = aの値と同じ
}