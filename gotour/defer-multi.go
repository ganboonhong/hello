package main

import "fmt"

func main(){
    // last-in-first-out
    defer fmt.Println(" world")
    defer fmt.Println("hello")
}