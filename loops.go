package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

// for d :=0;d<len(days);d++{
// 	fmt.Println(days[d])
// }

// for index,day :=range days{
// 	fmt.Printf("index is %v and value is %v\n",index,day)


	rougeValue :=1
	for rougeValue <10{

		if rougeValue ==5{
			break
		}
		fmt.Println("value is",rougeValue)
		rougeValue++
	}
}
      //}
