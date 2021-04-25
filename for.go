package main

import "fmt"

// メイン関数
func main() {
    // 普通のforの繰り返し
    for i := 0; i < 10; i++ {
        if i == 3 {
            //continue 処理を飛ばし繰り返し
            continue
        } else if i == 8 {
            break
        }

        fmt.Println(i)

    }

	// 多重ループの内側のfor文から外側のfor文を超えて抜け出る
	LOOP:
	for i := 0; i < 3; i++{
		for j := 0; j < 3; j++{
			if i == 1 && j == 2{
				break LOOP
			}
			fmt.Println(i, j)
		}
	}

    // while文的にも作成可能
    n := 0
    for n < 10 {
        fmt.Println(n)
        n++
    }
}