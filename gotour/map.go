package main

import "fmt"

type Food struct {
	Name  string
	Price int
}

type Car struct {
	Name  string
	Price int
}

var m map[int]Food

func main() {
	m = make(map[int]Food)
	m[123] = Food{
		"chichen rice", 5,
	}
	fmt.Println(m)

	m2 := map[int]Car{
		456: {
			"Toyota", 1000,
		},
	}

    fmt.Println(m2)
	fmt.Println(m2[456])
}