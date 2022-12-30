package main

import "fmt"

const LoginToken string = "asjaioquwj";

func main() {
	
	var message string = "hello from the other world"
	fmt.Println(message)
	fmt.Printf("Type of message is %T \n", message)

	var message2 = "hello from the 2nd world"
	fmt.Println(message2)
	fmt.Printf("Type of message2 is %T \n", message2)

	message3 := "hello from the 3rd world"
	fmt.Println(message3)
	fmt.Printf("Type of message3 is %T \n", message3)

	fmt.Println(LoginToken)
	fmt.Printf("Type of LoginToken is %T \n", LoginToken)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn is %T \n", isLoggedIn)

	var numberOfUsers uint8 = 255
	fmt.Println(numberOfUsers)
	fmt.Printf("Type of numberOfUsers is %T \n", numberOfUsers)

	var amount float32 = 456.102987129817
	fmt.Println(amount)
	fmt.Printf("Type of amount is %T \n", amount)

	// default values
	var color string
	fmt.Printf(color)

	var noOfWindows uint8
	fmt.Println(noOfWindows)

	var isChecked bool
	fmt.Println(isChecked)

	var length float32
	fmt.Println(length)
}