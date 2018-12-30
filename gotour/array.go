package main

import "fmt"

func main(){
    var a [3]string
    a[0] = "gan"
    a[1] = "boon"
    a[2] = "hong"

    fmt.Println(a[0])
    fmt.Println(a)

    sentence := [2]string{"learning", "Go"}
    fmt.Println(sentence)

}