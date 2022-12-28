package main

import "fmt"

func main() {
	fmt.Println("welcome to the go array class")

	var foodlist [4]string

	foodlist[0] = "Biryani"
	foodlist[1] = "Keema"
	foodlist[3] = "chicken"

	fmt.Println("fool list is : ", foodlist)
	fmt.Println("fool list is : ", len(foodlist))

	var fruitList = [3]string{"Banana", "Apple", "peach"}

	fmt.Println("The list fruit is : ", fruitList)
	fmt.Println("The list fruit is : ", len(fruitList))

}
