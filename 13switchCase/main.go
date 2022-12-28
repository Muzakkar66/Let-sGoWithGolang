package main

import (
	"fmt"
	"math/rand" //package
	"time"
)

func main() {
	fmt.Println("Welcome in switch and case in golang")
	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1 //it will generate btween 0 to 5 rang and then adding 1 into genrated number from 0 to 5
	fmt.Println("Value of dice is: ", diceNumber)

	// switch and case statment implementing
	//agar diceNumber 3 generate kary to 3 wala case chal jay
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 ")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 4")
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6 you can roll out again")
	default:
		fmt.Println("What was that!")
	}
}
