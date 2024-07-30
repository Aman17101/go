package main

import "fmt"
func main(){
	fmt.Println("welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0]="Apple"
	fruitlist[1]="Bananna"
	fruitlist[2]="peach"

	fmt.Println("fruitlist is :", fruitlist)
	fmt.Println("fruitlist length is : ",len(fruitlist))

	var vegList=[3]string{"potato" ,"Tomato" ,"Carrot"}
	fmt.Println("vegy list is :",len(vegList))
}
