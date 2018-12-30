package main

import(
    "fmt"
    "time"
)

func main() {

    today := time.Now().Weekday();
    sat := "Saturday"
    fmt.Println(today.String() == sat)
    // fmt.Println(int(today) == 6)
}