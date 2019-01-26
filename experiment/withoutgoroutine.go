package main

import (
    "fmt"
)

func countAsc(){
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }
}

func countDesc(){
    for i := 100; i >= 0; i-- {
        fmt.Println(i)
    }
}

func main(){
    countAsc();
    countDesc();
}