package main

import "fmt"

// Go has only one loop structure: for
func main() {
	// Example 1, Standard for loop
	sum := 0
	for i := 0; i < 20; i += 2 {
		sum += i
	}
	fmt.Println(sum) // 90

	// Example 2, Omitted
	// The init and post statements are optional.
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum) // 1024

	// Example 3, if is while in Go
	// At that point you can drop the semicolons:
	// C's while is spelled for in Go.
	sum = 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum) // 1024

	// Example 4, Forever
	// If you omit the loop condition it loops forever
	// so an infinite loop is compactly expressed.
	// output: timeout running program
	//for {
	//	fmt.Println(sum)
	//}
}
