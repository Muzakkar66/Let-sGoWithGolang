package main

import "fmt"

const JwtToken string = "DWRR3544t4" //public
func main() {
	var firstName string = "Azad Develper"
	fmt.Println(firstName)
	fmt.Printf("The type of verable is: %T \n", firstName) //so printf is printing and println also printing same kind of stuff

	var IsLogged bool = true
	fmt.Println(IsLogged)
	fmt.Printf("Type of: %T \n", IsLogged)

	var smalVal uint8 = 255
	fmt.Println(smalVal)
	fmt.Printf("Type of var: %T \n", smalVal)

	var smallFlot float64 = 234.444345353453454444
	fmt.Println(smallFlot)
	fmt.Printf("Type of var: %T \n", smallFlot)

	fmt.Println(JwtToken)
	fmt.Printf("Type of var: %T \n", JwtToken)

	//no var style
	noVarStyle := "hello"
	fmt.Println(noVarStyle)
	fmt.Printf("Type of Var: %T \n", noVarStyle)


}
