package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("welcome to the making get req in golang")

	const myUrl = "http://localhost:3000/users" //here i added nestjs running server url

	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Statues code: ", res.StatusCode)
	fmt.Println("Content length :", res.ContentLength)

	//now printing the execul content frome the nest
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println("The content body is : ", string(content))
}
