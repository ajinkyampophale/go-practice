package main

import "fmt"

func main() {

	fmt.Println("Loops")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Printf("Index: %v , Value: %v \n", d, days[d])
	// }

	// for d := range days {
	// 	fmt.Printf("Index: %v , Value: %v \n", d, days[d])
	// }

	// for index, value := range days {
	// 	fmt.Printf("Index: %v , Value: %v \n", index, value)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if(rougueValue == 2) {
			rougueValue++
			continue
		}

		if(rougueValue == 7) {
			break
		}

		if(rougueValue == 4) {
			goto gotoLabel
		}

		fmt.Println("Rougue Value", rougueValue)

		rougueValue++
	}

	gotoLabel:
		fmt.Println("Go to executed")

}