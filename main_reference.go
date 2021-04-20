package main

import (
	// f"fmt" 任意の名前指定可能
	"fmt"
)

func main() {
	// 変数の定義 ver 変数名 型
	var name string
	// 宣言の省略(型は自動的に決まる)
	names := サンプル君

	/* 定数の定義(グローバルとして使用する場合は関数外
	で定義する)*/
	const PI = 3.14
	// 自動で1づつ増やす
	const (
		Sun int = iota
		Mon
		Tue
	)

	// f.Printf("名前を入力してください") 任意の名前指定したver
	fmt.Printf("名前を入力してください")
	fmt.Scan(&name)
	fmt.Printf("こんにちは,%sさん!", name)
}
