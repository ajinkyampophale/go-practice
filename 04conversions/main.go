package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please rate our Pizza")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", rating);

	convertedRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if(err != nil) {
		fmt.Println("Conversion error ", err)
	} else {
		fmt.Println("Adding one to your rating ", convertedRating + 1)
	}

}