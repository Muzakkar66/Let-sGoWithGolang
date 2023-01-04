package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers in golang")
	var pts *int

	fmt.Println(pts)

	myNum := 23

	var ptr = &myNum
	fmt.Println("The Referance value of actual pointer is :", ptr)
	fmt.Println("The value of actual pointer is :", *ptr)
	*ptr = *ptr + 2

	fmt.Println("The the result of pointer is: ", myNum)
}
