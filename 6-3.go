package main

import (
	"fmt"
)

func enzan(a, b int) (int, int){
	add := a + b
	sub := a - b
	return add, sub
}

func main() {
	var a int
	var b int
	var add int
	var sub int
	fmt.Scan(&a, &b)
	add, sub = enzan(a, b)
	fmt.Println(add)
	fmt.Println(sub)
}