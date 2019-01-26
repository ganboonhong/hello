package main

import (
    "fmt"
)

func countDesc(){
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }
}

func countAsc(){
    for i := 100; i >= 0; i-- {
        fmt.Println(i)
    }
}

func main(){
    go countDesc();
    countAsc();
}