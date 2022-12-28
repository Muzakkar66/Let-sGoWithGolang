//map in golang mean you can map values with key pairs like JS equl to JAvaScript ....key value pairs mean
package main

import (
	"fmt"
)

func main() {

	fmt.Println("welcome to the maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["go"] = "Golang"

	fmt.Println(languages)       // in result you getting key values pairs in map
	fmt.Println(languages["go"]) //getting value through key

	//let remove some value through key
	delete(languages, "RB")
	fmt.Println(languages) //in result will Ruby

	//loops are interesting in golong
	for key, value := range languages {
		fmt.Println("for key v, value is........... %v  ", key, value)
	}

	//loops are interesting in golong
	for _, value := range languages {
		fmt.Println("for key v, value is %v  ", value)
	}

}
