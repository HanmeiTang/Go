package main

import "fmt"

func main() {
	var s1 = []int{}
	isNil(s1)

	var s2 []int
	isNil(s2)
}

func isNil(i []int) {
	if i == nil {
		fmt.Printf("This is %v, which is a nil!\n", i)
	} else {
		fmt.Printf("This is %v, which is NOT a nil!\n", i)
		fmt.Printf("len: %d; cap: %d; val: %v\n", len(i), cap(i), i)
	}
}
