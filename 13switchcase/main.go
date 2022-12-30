package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")

	rand.Seed(time.Now().UnixNano())
	var dice = rand.Intn(6) + 1

	switch dice {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	case 3:
		fmt.Println("Case 3")
	case 4:
		fmt.Println("Case 4")
		fallthrough
	case 5:
		fmt.Println("Case 5")
	case 6:
		fmt.Println("Case 6")
	default:
		fmt.Println("Default")
	}
}