package main

import (
	"fmt"
)

func wasa(a, b int) (int, int){
	wa := a + b
	sa := a - b
	if sa < 0 {
		sa *= -1
	}
	return wa, sa
}

func main() {
	var wa, sa int
	wa, sa = wasa(5,2)
	fmt.Printf("和=%d, 差=%d\n", wa, sa)
}