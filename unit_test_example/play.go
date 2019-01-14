package main

import (
    // "fmt"
    "log"
    // "os/exec"
)

func Calc(n int) int{
    return n + 1
}

func main() {
    // cmd := exec.Command("ls", "-lah")
    // cmd := exec.Command("say", "hello")
    // out, err := cmd.CombinedOutput()
    // if err != nil {
    //     log.Fatalf("cmd.Run() failed with %s\n", err)
    // }
    // fmt.Printf("combined out:\n%s\n", string(out))
    log.Println(Calc(2))
}