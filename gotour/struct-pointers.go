package main

import "fmt"

type Obj struct {
    X int
    Y int
}

func main(){
    o := Obj{1, 2}
    p := &o
    p.X = 3
    fmt.Println(o)
}