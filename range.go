package main

import "fmt"

// メイン関数
func main() {
    // スライス
    s := []int{2, 3, 8}

    for i, v := range s {
        fmt.Println(i, v)
    }

    // ブランク修飾子
    // _にしておくと何が入ってきてもそれを廃棄してくれるという仕様
    for _, v := range s {
        fmt.Println(v)
    }

    // マップ
    m := map[string]int{"fujimoto": 200, "arita": 300}

    for k, v := range m {
        fmt.Println(k, v)
    }
}