package main

import "fmt"

func main() {
	var aa [2]string
	aa[0] = "Hello"
	aa[1] = "World"

	fmt.Println(aa[0], aa[1]) // Hello World
	fmt.Println(aa)           // [Hello World]

	primes := [10]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13 0 0 0 0]

}
