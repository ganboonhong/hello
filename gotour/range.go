package main

import "fmt"

func main(){
    alphabets := []string{"a", "b", "c"}

    for i, v := range alphabets {
        fmt.Println("index: ", i, "value: ", v);
    }
}