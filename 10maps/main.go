package main

import "fmt"

func main() {
	fmt.Println("Maps")

	var languages = make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages", languages)
	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Languages", languages)

	// loops
	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}

	for _, value := range languages {
		fmt.Println("Value:", value)
	}

}
