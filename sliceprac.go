// スライスは基本的に配列と同期しているが配列の容量を超えるとメモリーに新しく領域が確保され配列と同期しなくなる
package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	var slc = arr[0:2]
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("arr=", arr)
	fmt.Println("slc=", slc)
	fmt.Printf("len=%d,cap=%d \n",len(slc), cap(slc))
	slc = append(slc, 50) //スライスに3番目の要素追加
	//この時点で、slc変数が別メモリに再定義しなおされる
	//arrとslcは連携しなくなる
	fmt.Println("arr=", arr)
	fmt.Println("slc=", slc)
	slc = append(slc, 100) //スライスに4番目の要素追加->容量を超えている
	fmt.Println("arr=", arr)
	fmt.Println("slc=", slc)
	fmt.Printf("len=%d,cap=%d \n",len(slc), cap(slc))
	slc[1] = -20 //スライスの値を変えても、配列の値は変わらない
	fmt.Println("arr=", arr)
	fmt.Println("slc=", slc)
}