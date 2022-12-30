package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[3] = "Banana"

	fmt.Println("Fruit List", fruits)
	fmt.Println("Fruit List Length", len(fruits))

	var vegList = [5]string{"potato", "beans", "spinach"}

	fmt.Println("Veg List", vegList)
	fmt.Println("Veg List Length", len(vegList))

	
}