package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://muhammadmuzakkarazad.com/about?experiece=4year&portfolio=2projectsDone"

func main() {
	fmt.Println("welcome to handing url in golang")

	fmt.Println(myurl)

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)

	fmt.Println(result.Path)
	fmt.Println(result.Port())

	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("The type of query params are:%T", qparams)
	fmt.Println(qparams["courses"])// here the mapnig the query params in golang

}
