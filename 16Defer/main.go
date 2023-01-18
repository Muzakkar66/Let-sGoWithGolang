package main

import "fmt"

func main() {
	defer fmt.Println("Hello defer ......u duffer actually 1 ")
	defer fmt.Println("Hello defer ......u duffer actually 2")
	defer fmt.Println("Hello defer ......u duffer actually 3")
	defer fmt.Println("Hello defer ......u duffer actually 4\n")

	fmt.Println("Welcome to the defer in golang \n")
	defer fmt.Println("Hello defer ......u duffer actually \n") //mean last in first out

	mydefer()
}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)

	}
}
