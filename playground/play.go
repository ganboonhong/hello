package main

import (
    "fmt"
)

func count0to1k(){
    for i := 1; i <= 100; i++ {
        fmt.Println(i)
    }
}

func count1kto0(){
    for i := 100; i >= 0; i-- {
        fmt.Println(i)
    }
}

func main(){
    go count0to1k();
    count1kto0();
}