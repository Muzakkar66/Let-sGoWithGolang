package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the our Pizza App!")
	fmt.Printf("Please rate out Pizza between 1 to 5 \n")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

}
