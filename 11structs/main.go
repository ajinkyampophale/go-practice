package main

import "fmt"

func main() {
	fmt.Println("Structs")

	aj := User{"Ajinkya", "aj@go.dev", true, 16}
	fmt.Println(aj)
	fmt.Printf("Formatted Details %+v \n", aj)
	fmt.Printf("Name: %v", aj.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}