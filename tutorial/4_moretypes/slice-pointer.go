package main

import "fmt"

func main() {
	var names [4]string = [4]string{
		"James",
		"David",
		"Smith",
		"Harry",
	}
	a := names[0:1]
	b := names[2:4]

	fmt.Println(a, b)

	a[0] = "Hanmei"
	fmt.Println(names)
}
