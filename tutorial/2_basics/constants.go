package main

import "fmt"

const Pi = 3.14

func main() {
	// Similar to var variable declaration
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Hello", Pi, "Day!")

	// constant variable can not be declared using :=
	const Truth = true
	fmt.Println("Go rules?", Truth)

}
