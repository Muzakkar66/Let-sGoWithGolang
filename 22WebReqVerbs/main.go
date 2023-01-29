package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to the wen request for post form data")

	// PostRequestAPI()
	PerformFormData() // but its not works becs in nest.js store form data route is not added properly

}

// func PostRequestAPI() {

// 	const myWeb = "http://localhost:3000/users"

// 	//fake json payload
// 	requestBody := strings.NewReader(`
// 	{
// 		"name":"Muhammad Muzakkar Azad",
// 		"age":100

// 	}
// 	`) // using ` ` back ticks
// 	response, err := http.Post(myWeb, "application/json", requestBody)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	responseOfPostReq, _ := ioutil.ReadAll(response.Body)

// 	fmt.Println(string(responseOfPostReq))
// }

func PerformFormData() {
	const myUrl = "http://localhot:3000/users"

	//formdata
	data := url.Values{}
	data.Add("firstname", "Muzakkir")
	data.Add("lastname", "Azad")
	data.Add("email", "muzakkir066@gmail.com")
	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
