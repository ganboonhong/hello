package main

import "fmt"

func main(){
    i, j := 10, 20
    p := &i
    fmt.Println(*p)
    *p = *p * j
    fmt.Println(*p)
    fmt.Println(p)
}