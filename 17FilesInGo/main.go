package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in Golong")

	content := "muhammad-muzakkir.asad@ceative.co.uk"
	file, err := os.Create("./mylcofile.txt") //here OS is package or module

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	 
	length, err := io.WriteString(file, content) // here io is a package
	checkNilErr(err)
	fmt.Println("length is :", length)
	defer file.Close() //here defer key word putting mean we will writing some code in below of closing
	readFile("./mylcofile.txt")

}

func readFile(fileName string) {
	dataBytes, err := ioutil.ReadFile(fileName) //ioutil is saprite utility for read a file
	checkNilErr(err)
	fmt.Println("After reading file result is: ", string(dataBytes))

}
func checkNilErr(err error) { //using above for save the lines of code
	if err != nil {
		panic(err)
	}
}
