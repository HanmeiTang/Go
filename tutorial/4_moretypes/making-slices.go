package main

import "fmt"

func main() {
	a := make([]int, 3) // len(a)=5
	printSlice2("a", a)

	b := make([]int, 5, 15)
	printSlice2("b", b)

	c := b[:15]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
