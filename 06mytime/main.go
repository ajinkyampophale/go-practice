package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Time study")

	presentTime := time.Now()
	fmt.Println("Present Time", presentTime)
	fmt.Printf("Type of time: %T", presentTime)

	fmt.Println("Formated Date:", presentTime.Format("01-02-2006"))
	fmt.Println("Formated Date and Day:", presentTime.Format("01-02-2006 Monday"))
	fmt.Println("Formated Date and Day and Time:", presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println("Formated Date and Day and Time OR:", presentTime.Format("01-02-2006 Monday 15:04:05"))

	customDate := time.Date(2000, time.August, 10, 4, 23, 0, 0, time.UTC)
	fmt.Println("Custom Date:", customDate)
	fmt.Println("Custom Date Formatted:", customDate.Format("01-02-2006 Monday"))
}