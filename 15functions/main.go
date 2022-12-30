package main

import "fmt"

func main() {
	fmt.Println("Functions")

	greetings();

	result := sum(5, 6)
	fmt.Println("Sum", result)

	proResult, message := proSum(4, 5, 6, 9)
	fmt.Println("Pro Result", proResult)
	fmt.Println("Pro Message", message)
}

func greetings() {
	fmt.Println("Hello from greetings")
}

func sum(a int, b int) int {
	return a + b
}

func proSum(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}

	return total, "Hello returning string"
}