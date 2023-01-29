package main


import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to making post request in golang")
	PostRequestAPI()
}

func PostRequestAPI() {

	const myWeb = "http://localhost:3000/users"

	//fake json payload
	requestBody := strings.NewReader(`
	{
		"name":"Muhammad Muzakkar Azad",
		"age":100
		
	}
	`) // using ` ` back ticks
	response, err := http.Post(myWeb, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	responseOfPostReq, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(responseOfPostReq))
}
