package main

import "fmt"
const LogInToken string = "jdkfiweu4583485748y78" //its public varible because variable defined in first capital latter

func main() {

	var username string = "@@@@@@@@@@@@@@  Muhammad Muzakkir Azad @@@@@@@@@@@@@@@@@@"
	fmt.Println(username)
	fmt.Printf("Variables type is: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables type is: %T \t\t\t\t\t\t", isLoggedIn) // \t for spacing

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("variables type is : %T \n", smallValue)
	 
	var floatValue float32= 225.534543634533 //in 32 5 number will show after . 5 no
	fmt.Println(floatValue) 
	fmt.Printf("variables type is : %T \n", floatValue)

	//defualt value and some aliases
	var anotherVariables int
	fmt.Println(anotherVariables)
	fmt.Printf("variables type is : %T \n", anotherVariables)

	//implicit type
	var website = "muhammmadmuzakkarazad.com" //if you not defining the type of string then g auto set according to variable value  
	fmt.Println(website)
	fmt.Printf("Type of varible is: %T \n", website)

	//no var stype
	numberOfVariable := 300000.0 // For example, foo := 32 is a short-hand form of var foo int mean not add var keyword and type 
	fmt.Println(numberOfVariable)

	//printing public variable as above declar
	fmt.Println(LogInToken)
	fmt.Printf("Type of varible is: %T \n", LogInToken)



}
