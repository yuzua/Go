package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	//os.Stdinでキーボードからの入力を受け付ける
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("%d",name)
	{

	}
}