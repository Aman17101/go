package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("welcome to handling url in golang")
	fmt.Println(myurl)

//parsing
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)                 
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams :=result.Query()
	fmt.Printf("The type of query params are :%T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _, val :=range qparams{
		fmt.Println("Param is :",val)

		partsofUrl :=&url.URL{
			Scheme: "https",
			Host: "lco.dev",
			Path: "/learn",
			RawPath: "user=Aman",
			
		}
		anotherURL :=partsofUrl.String()
		fmt.Println(anotherURL)
	}
}
