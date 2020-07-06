package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q) // [1 2 3 4 5]

	r := []bool{false, true, true, true, false}
	fmt.Println(r) // [false true true true false]

	s := []struct {
		X int
		Y string
	}{
		{1, "Hanmei"},
		{2, "Harry"},
		{3, "Hanker"},
	}
	fmt.Println(s) // [{1 Hanmei} {2 Harry} {3 Hanker}]
	fmt.Println(s[1:2])
	fmt.Println(s[1:])
	// fmt.Println(s[1:5]) Error
}
