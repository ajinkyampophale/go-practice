package main

import "fmt"

func main() {
	fmt.Println("Structs")

	aj := User{"Ajinkya", "aj@go.dev", true, 16}
	fmt.Println(aj)
	aj.GetStatus()
	aj.NewMail()
	fmt.Println("Mail Id after malipulation:", aj.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User is active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New Mail Id:", u.Email)
}