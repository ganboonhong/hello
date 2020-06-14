package reference

import (
    "testing"
)

// https://tutorialedge.net/golang/intro-testing-in-go/

func TestCalc(t *testing.T){
    if Calc(2) != 3 {
        t.Error("Expected 2 + 1 to equal 3")
    }
}


// Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the “go test” command, which automates execution of any function of the form
// func TestXxx(*testing.T)
// where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.
// https://golang.org/pkg/testing/ > Overview

func TestMultiple(t *testing.T){
    for _, test := range []struct{
        input int
        expected int
    }{
        {1, 2},
        {2, 3},
        {3, 4},
    }{
        if output := Calc(test.input); output != test.expected {
            t.Errorf("Test failed: %d input, %d expected, output is %d", test.input, test.expected, output)
        }
    }
}