package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcom string = " Welcome to the user input" // Note: It is not possible to declare a variable using ":=" without assigning a value to it.
	welcome := welcom // no var style
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for Pizza")

	// comma syntax or err ok
	input, _ := reader.ReadString('\n')
	fmt.Println(" Thanks for rating ", input)
	fmt.Printf("Type of string %T", input)



	// welcomeVar := welcom1
	// fmt.Println(welcomeVar)

}
 