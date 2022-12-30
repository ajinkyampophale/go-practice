package main

import "fmt"

func main() {

	var ptr *int
	fmt.Println("Pointer values:", ptr)

	myNumber := 23

	var numPtr = &myNumber

	fmt.Println("Num Reference Pointer:", numPtr)
	fmt.Println("Num Reference Pointer Value:", *numPtr)

	*numPtr = *numPtr + 2;

	fmt.Println("Og Value:", myNumber);

}