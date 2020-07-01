package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	//fmt.Print(today)
	gap := int(time.Saturday - today)

	switch {
	case gap == 0:
		fmt.Println("Today")
	case gap == 1:
		fmt.Println("Tomorrow")
	case gap == 2:
		fmt.Println("In 2 days")
	default:
		fmt.Printf("Too far away... %v days left.\n", gap)
	}

}
