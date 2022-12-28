package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the input user")
	fmt.Println("Please enter the Pizza rating:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating God Bless you: ", input)

	//adding 1 in input rating to convert string float 

	numRating, err:= strconv.ParseFloat(strings.TrimSpace(input), 64)

	//like try catch
	if err != nil {
		fmt.Println(err)
	} else {
		println("Added 1 to your rating: ", numRating+1)
	}

}
	