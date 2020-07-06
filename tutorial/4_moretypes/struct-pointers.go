package main

import "fmt"

// Create a new type named 'Vertex2' based on `struct`
type Vertex2 struct {
	X int
	Y int
}

func main() {
	v := Vertex2{1, 2}
	fmt.Println(v.X) // 1
	v.X = 5
	fmt.Println(v.X) // 5
	p := &v
	fmt.Println(p)  // &{5 2}
	fmt.Println(*p) // {5 2}
	fmt.Println(&p) // 0xc00000e030
	fmt.Println(&v) // &{5 2}

	p.X = 1e9
	fmt.Println(v.X) // 1000000000
	(*p).X = 1
	fmt.Println(v.X) // 1

}
