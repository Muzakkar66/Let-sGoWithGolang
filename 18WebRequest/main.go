package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Welcome to in web request in golang")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("response is of type:%T\n", response)
	defer response.Body.Close() // coller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(databyte)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
