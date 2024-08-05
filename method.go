package main

import "fmt"

func main() {
	fmt.Println("welcome to method in golang")

	//no inheritence  in golang ;No super or Parent

	Aman := User{"Aman", "aman@gov.in", true, 19}
	fmt.Println(Aman)
	fmt.Printf("Aman details are: %+v\n", Aman)
	fmt.Printf("Name is %v and email is %v.\n", Aman.Name, Aman.Email)
	Aman.GetStatus()
	Aman.NewMail()
	fmt.Printf("Name is %v and email is %v \n", Aman.Name, Aman.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New mail is created for user:", u.Email)
}
