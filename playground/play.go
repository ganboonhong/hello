package main

import (
    "fmt"
    "os"
)

func main(){
    token := os.Getenv("GO_PATH")
    fmt.Println(token)
}