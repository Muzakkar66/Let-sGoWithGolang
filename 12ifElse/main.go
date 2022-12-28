package main

import "fmt"

func main() {
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Not regular User"
	} else {
		result = "Exectly 10 login count"
	}

	//secod example 
	if 9%2 == 0 {
		fmt.Println("Number is even")
	}else {
		fmt.Println("Number is odd")
	}

	if num:= 10; num >10{
		fmt.Println("Number is greater then 10")
	} else{
		fmt.Println("Number is less then 10")
	}

	fmt.Println(result)
}
