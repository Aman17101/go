package main



import "fmt"

const LoginToken string ="uegjsjiojiogdf" //public

func main(){
	var username string= "Aman"
	fmt.Println(username)
	fmt.Printf("variable of type : %T \n",username)

	var isLoggedIn bool= false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable of type : %T \n",isLoggedIn)

	var smallVal uint=255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n ",smallVal)

	var floatVal float64=255.546382202738
	fmt.Println(floatVal)
	fmt.Printf("variable is of type: %T \n ",floatVal)

	//default values of some aliases
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type :%T \n",anotherVariable)

	//implicit type
	var implicitVar = 1000
	fmt.Println(implicitVar)

	//no var style

	numberOfUser :=30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n ",LoginToken)



}
