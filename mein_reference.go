package main

import (
	// f"fmt" 任意の名前指定可能
	"fmt"
)

func main() {
	// 変数の定義 ver 変数名 型
	var name string
	// f.Printf("名前を入力してください") 任意の名前指定したver
	fmt.Printf("名前を入力してください")
	fmt.Scan(&name)
	fmt.Printf("こんにちは,%sさん!", name)
}
