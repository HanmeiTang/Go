package main

import (
	"fmt"
)

const (
	// Big = 2^100
	Big = 1 << 100
	// Small = 2^1 = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// constant 1267650600228229401496703205376 overflows int
	fmt.Println(needInt(Big))
}
