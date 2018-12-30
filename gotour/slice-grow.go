package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = s[2:5]

	fmt.Println("len: ", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))

	s[2] = 6

	fmt.Println("len: ", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))

	t := make([]int, len(s), (cap(s)+1)*2)
	copy(t, s)

    t[2] = 7
	fmt.Println("len: ", t)
	fmt.Println("len: ", len(t))
	fmt.Println("cap: ", cap(t))
}
