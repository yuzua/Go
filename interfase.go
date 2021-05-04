package main

import (
	"fmt"
)

func main() {
	// 入れた値によって勝手に型を変更する(Pythonみたい)
	 var i interface{}
	 i = 100
	 fmt.Printf("%T\n",i) // int
	 i = 0.5
	 fmt.Printf("%T\n",i) // float64
	 i = "hello"
	 fmt.Printf("%T\n",i) // string
	 // interface型の調べ方 i.(type)
	 // 使い方
	switch i.(type) {
	case string:
		fmt.Printf("文字列型です")
	}
}