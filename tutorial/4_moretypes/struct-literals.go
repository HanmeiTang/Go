package main

import "fmt"

type Vetex3 struct {
	X, Y int
}

var (
	v1 = Vetex3{1, 2}
	v2 = Vetex3{X: 3}
	v3 = Vetex3{}
	p  = &Vetex3{4, 5}
)

func main() {
	// {1 2} {3 0} {0 0} &{4 5}
	fmt.Println(v1, v2, v3, p)
}
