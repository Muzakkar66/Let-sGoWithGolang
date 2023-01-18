package main

import "fmt"

func main() {

	fmt.Println("welcome to the methods")

	structt := User{"Muhammad Muzakkir", "muzakkir066@gmail.com", true, 26}
	fmt.Println("Printing struct value______________", structt.Name, structt.Email, structt.Status, structt.Age)

	structt.GetNameFromStruct()

	structt.EmailReplacing() //here calling method after

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//here is the method
func (u User) GetNameFromStruct() { //u is the object of struct user for using
	fmt.Println("The name from struct properties:", u.Name)
}

//now manuplating the email from string in slice use method

func (e User) EmailReplacing() {
	e.Email = "muhammad-muzakkir.asad@ceative.co.uk"
	fmt.Println("The Email from the struct after manipulating: ", e.Email)
}
