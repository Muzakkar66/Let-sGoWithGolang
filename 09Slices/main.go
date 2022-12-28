package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Slice in golang")

	var fruitList = []string{"Banana", "Apple", "Peach"}
	fmt.Println("The type of obove variable is %T\n", fruitList)

	fruitList = append(fruitList, "Orange", "Water Malen") //appending mean adding value into the array
	fmt.Println("After appanding some values: ", fruitList)

	fruitList = append(fruitList[1:]) //its mean start from arry index 1
	fmt.Println("After putting index 1 from to print out:", fruitList)

	fruitList = append(fruitList[1:3]) //its mean start from arry index 1 to 3  mean the result will be 1,2 indexes
	fmt.Println("After putting from index 1  to 3 printing out:", fruitList)

	fruitList = append(fruitList[:3]) // mean start from 0 to 3 in result 3 also incusive
	fmt.Println("After appending the result: ", fruitList)

	highScore := make([]int, 4) //[]int is a slice
	highScore[0] = 231
	highScore[1] = 232
	highScore[2] = 230
	highScore[3] = 237

	highScore = append(highScore, 122, 443, 232)
	fmt.Println(highScore)

	//sorting slice
	sort.Ints(highScore)

	fmt.Println("Sorting slice: ", highScore)
	fmt.Println(sort.IntsAreSorted(highScore)) // checking slice is sorted or not through sort function

	//removing index value in slice

	var courses = []string{"Reactjs", "Python", "Ruby", "NodeJs"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //mean removing index 2 from the slice
	fmt.Println(courses)
}
