package main

import (
    "testing"
)

func TestCalc(t *testing.T){
    if Calc(2) != 3 {
        t.Error("Expected 2 + 1 to equal 3")
    }
}