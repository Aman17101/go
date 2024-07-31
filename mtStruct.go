package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritence  in golang ;No super or Parent

	Aman := User{"Aman", "aman@gov.in", true, 19}
	fmt.Println(Aman)
	fmt.Printf("Aman details are: %+v\n", Aman)
	fmt.Printf("Name is %v and email is %v.", Aman.Name, Aman.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
