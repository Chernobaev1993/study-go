package main

import "fmt"

func main() {
    fmt.Println("начало")

    for i := 0; i < 5; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("конец")
}
