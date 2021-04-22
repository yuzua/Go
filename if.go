package main

import "fmt"

// メイン関数
func main() {
    // 普通の条件分岐
    score := 61

    if score > 80 {
        fmt.Println("Great!")
    } else if score > 60 {
        fmt.Println("Good!")
    } else {
        fmt.Println("soso")
    }

    // golangの特徴の条件分岐
    if score := 52; score > 80 {
        fmt.Println("Great!")
    } else if score > 60 {
        fmt.Println("Good!")
    } else {
        fmt.Println("soso")
    }
}