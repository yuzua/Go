// hello.go
package main

import (
	"fmt"
)

func byVal(a int){
	a = a+1
}

func byRef(a *int){
	*a = *a +1
	// aのアドレスを調べるには"%p"を使用
	fmt.Printf("%p\n", a)
}

func main() {
	var a int
	var p *int
	a = 100
	p = &a // "&"はアドレス演算子 "a"のアドレスを調べて,その値を"p"に代入
	fmt.Printf(*p) // "*"はアドレスの中身 = aの値と同じ

	var n = 1
	byVal(n)
	fmt.Println("n=",n)

	n = 1
	byRef(&n)
	fmt.Println("n=",n)
}