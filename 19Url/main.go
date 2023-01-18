package main

import "fmt"

const myurl string = "http://muhammadmuzakkarazad.com/about?experiece=4year&portfolio=2projectsDone"

func main() {
	fmt.Println("welcome to handing url in golang")

	fmt.Println(myurl)

	result, _ := myurl

	fmt.Println(result.Scheme)
}
