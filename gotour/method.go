package main

import "fmt"

type Word2Phrase struct {
    a, b string
}

func (p Word2Phrase) CatWord2Phrase() string {
    return p.a + " " + p.b
}

func main(){
    words := Word2Phrase{"before", "lunch"}
    phrase := words.CatWord2Phrase();
    fmt.Println(phrase)
}