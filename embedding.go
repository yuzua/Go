package main

import (
	"fmt"
)

type Point struct{
	X, Y int
}

type Circle struct{
	Point //JAVAの継承みたいなもの
	int //構造体の埋め込み
}

func main() {
	var c = Circle{Point{50,50}, 70}

	fmt.Println(c)
	fmt.Printf("中心座標は(%d,%d)\n", c.X, c.Y)
	fmt.Printf("半径は%d\n", c.int)
}