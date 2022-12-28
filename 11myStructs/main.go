package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")

	muzakkir := User{"Muhammad Muzakkir Azad", "muzakkir066@gmail.com", true, 26} //use struct to initilize the value

	fmt.Printf("struct output with details : %+v\n", muzakkir)
	fmt.Printf("Name is: %v and email is: %v Status is: %v\n", muzakkir.Name, muzakkir.Email, muzakkir.Status)

}

//making struck in golang
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
