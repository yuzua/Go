package main

import "fmt"

// メイン関数
func main() {
    signal := "blue"

    switch signal {
    case "red":
        fmt.Println("Stop")
    case "yellow":
        fmt.Println("caution")
    case "green", "blue":
        fmt.Println("Go")
    default:
        fmt.Println("wrong")
    }

    // if-else的にも書く事が可能
    score := 82
    switch {
    case score > 80:
        fmt.Println("Great!")
		// fallthroughは次のcaseかdefaultも実行できるようにする
		fallthrough
    default:
        fmt.Println("Bad")
    }
}