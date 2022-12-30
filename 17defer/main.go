package main

import "fmt"

func main() {

	// defer => Moves to End of function and is Exexuted in LIFO fashion

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	Defer()
	
}

func Defer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}