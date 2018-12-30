package main

import "fmt"

type Wordpad interface {
	Typing(s string) string
}

func (s MyString) Typing(echo string) string {
    return echo
}

type MyString string

func main() {
    var w Wordpad

    u := MyString("str")
    w = u
    fmt.Println(w.Typing("whatever you typed will be echoed"));

}