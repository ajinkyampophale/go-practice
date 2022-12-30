package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {

	count := 5
	count2 := 4.5

	fmt.Println("Total count ", count+int(count2))

	// random using math
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random", rand.Intn(5) + 1)

	// random using crypto
	myRandomNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random", myRandomNo)

}