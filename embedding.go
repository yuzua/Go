package main

import (
	"fmt"
)

type Point struct{
	X, Y int
}

type Circle struct{
	// Center Point
	Point // JAVAの継承みたいなもの
	// Radius int
	int // 構造体の埋め込み
}

func main() {
	var c = Circle{Point{50,50}, 70}

	fmt.Println(c)
	// fmt.Printf("中心座標は(%d,%d)\n", c.Center.X, c.Center.Y)
	fmt.Printf("中心座標は(%d,%d)\n", c.X, c.Y)
	// fmt.Printf("半径は%d\n", c.Radius)
	fmt.Printf("半径は%d\n", c.int)
}