package main

import "fmt"

func main() {
	var s = "こんにちは、Go言語"
	d := s[0:6]
	e := s[9:20]
	fmt.Println(d)
	fmt.Println(e)
}