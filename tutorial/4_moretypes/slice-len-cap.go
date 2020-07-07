package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	// len: 6; cap: 6; value: [2 3 5 7 11 13]
	printSlice(s)

	s = s[:0]
	// len: 0; cap: 6; value: []
	printSlice(s)

	s = s[:4]
	// len: 4; cap: 6; value: [2 3 5 7]
	printSlice(s)

	s = s[2:]
	// len: 2; cap: 4; value: [5 7]
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len: %d; cap: %d; value: %v\n",
		len(s), cap(s), s)
}
