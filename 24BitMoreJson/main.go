package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"_"`
	Tags     []string `json:"tags,omitempty"` //here tags is slice in go for multiple store the tag like in array
}

func main() {
	fmt.Println("welcome to the bit more json")
	EncodeJson()
}
func EncodeJson() {
	LcoCourses := []course{
		{
			"ReactJS Bootcamp",
			300,
			"Udemy",
			"pas1234",
			[]string{"web-Dev", "JS"},
		},
		{
			"NestJS Bootcamp",
			350,
			"Udemy",
			"pas1234",
			[]string{"web-Dev", "JS", "NEST"},
		},
		{
			"NodeJS Bootcamp",
			330,
			"Udemy",
			"pas1234",
			nil,
		},
	}
	//package data as json data
	finalJson, err := json.MarshalIndent(LcoCourses, "prefix", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("FInal Encoded Json is: \n", string(finalJson))

}
