package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Premium user"
	} else {
		result = "Admin"
	}

	fmt.Println(result)

	if(9%2==0){
		fmt.Println("Number is even")
	}else{
		fmt.Println("Number is odd")
	}

	if num:=3;num<10{
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is not less than 10")
	}
}
