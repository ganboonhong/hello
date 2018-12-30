package main

import (
    "fmt"
    "math"
    "reflect"
)

func main() {
    x, y := 3, 4
    z := math.Sqrt(float64(x*x + y*y));
    fmt.Println(z)
    fmt.Println(reflect.TypeOf(z))
}