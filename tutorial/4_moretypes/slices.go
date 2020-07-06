package main

import "fmt"

func main() {
	var primes2 = [10]int{1, 3, 5, 7, 11, 13}
	var slice []int = primes2[1:4]
	fmt.Println(slice) // [3 5 7]
}
