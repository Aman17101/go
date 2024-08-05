package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")
	greeter()
	//greeterTwo()

	result := adder(3, 5)

	fmt.Println("result is :", result)

	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("pro result is :", proRes)
	fmt.Println("pro message is :", myMessage)
}

func adder(valOne int, vaTwo int) int {
	return valOne + vaTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function "
}

// func greeterTwo() {
// fmt.Println("Another Method")

// }
func greeter() {
	fmt.Println("Namstey from golang")
}
