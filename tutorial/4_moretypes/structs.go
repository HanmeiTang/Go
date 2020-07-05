package main

import "fmt"

// struct is a collection of fields
type Vertex struct {
	X, Y, m int
}

func main() {
	v := Vertex{1, 2, 3}
	fmt.Println(v)
	fmt.Println(v.X)
	v.X = 100
	fmt.Println(v.X)
}
