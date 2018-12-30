package main

import "fmt"

func main() {
    sum, i := 0, 1
    // for i := 0; i < 10; i++ {
    //     sum += i
    // }
    for ; sum < 10; {
        sum += i
        fmt.Println(i, sum)
    }
}