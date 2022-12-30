package main

import "fmt"

func main() {
	fmt.Println("If Else")

	loginCount := 10

	if loginCount < 10 {
		fmt.Println("Less than 10")
	} else if loginCount > 10 {
		fmt.Println("Greater than 10")
	} else {
		fmt.Println("Equal to 10")
	}

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 9 {
		fmt.Println("Less than 9")
	} else {
		fmt.Println("Greater than 9")
	}

	// if(err != nil){

	// }

}