package main

import "fmt"

func main() {
	fmt.Println("welome to the Loops")

	days := []string{"Sunday", "monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} //its slice
	fmt.Println(days)

	//looping the slice of days .... type one
	for d := 0; d < len(days); d++ { // mean    d ki length jab tak days sy less hai to  loop run hota rhy
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^", days[d]) //here d in [] mean 0 sy length jitni b hai wo value print kr dy
	}
	fmt.Println("$$$$$$$$$$ second for loop $$$$$$$$$$")
	//second type
	for i := range days { //here range mean its loop through any array or slice like for loop
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println("The index of days slice with value are: ", index, day)
	}
	//here you can remove index or i what ever u use in
	for _, day := range days { //here _ mean comma ok syntex
		fmt.Println("The days slice value are: ", day)
	}

	
	//
	rougValue := 1
	//break it in for loop
	for rougValue < 10 {
		if rougValue == 5 {
			break
		}
		fmt.Println("Value are: ", rougValue)
		rougValue++
	}
	//continue  it in for loop
	for rougValue < 10 {
		if rougValue == 2 {
			goto lco //lco is reserved word like if, for and much more
		}
		// if rougValue == 5 {
		// 	rougValue++
		// 	continue
		// }
		fmt.Println("Value aressss: ", rougValue)
		rougValue++
	}

	// here is goto

lco:
	fmt.Println("jumping on muhammadmuzakkirazad.com")
}
