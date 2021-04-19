package main

import "fmt"

func main() {
	var name string
	fmt.Printf("名前を入力してください:")
	fmt.Scan(&name)
	fmt.Printf("こんにちは,%sさん!", name)
}
