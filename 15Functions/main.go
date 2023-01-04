package main

import "fmt"

func main() {
	fmt.Println("Welcome to the function in golang")
	greeter()
	result := adder(3, 6) //two arguments
	fmt.Println("Result is : ", result)

	proResult, myMassage := proAdder(2, 4, 6, 7, 8)

	fmt.Println("The pro result is:", proResult)
	fmt.Println("The massage is :", myMassage)
}
func adder(valOne int, valTwo int) int { // before {} baraces add ...
	//function signature in  to adentifying the arguments that you passing out is type is this for passing and returning value on function
	return valOne + valTwo
}

//adding pro function
func proAdder(values ...int) (int, string) { // here three dots ... is paradic functions
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "the massage with signature text"
}
func greeter() {
	fmt.Println("Assalam u Alaikum golang learners*******")
}
