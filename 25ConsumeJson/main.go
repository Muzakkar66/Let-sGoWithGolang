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
	fmt.Println("welcome to the consume  json")
	// EncodeJson()
	DecodeJson()
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
	fmt.Println("Final Encoded Json is: \n", string(finalJson))

}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"courseName": "NestJS Bootcamp",
				  "Price": 350,
				  "website": "Udemy",
				  "_": "pas1234",
				  "tags": [
						  "web-Dev",
						  "JS",
						  "NEST"
				  ]
	}
	`)
	var LcoCourse course
	chackValid := json.Valid(jsonDataFromWeb)
	//chacking json is valid or not
	if chackValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &LcoCourse) //json ko unmarshal karna mean intent to un intent
		fmt.Printf("%#v\n", LcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some case you just want to add some data to key value

	var myOnlineMap map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineMap)
	fmt.Printf("%#v\n", myOnlineMap)

	for k, v := range myOnlineMap {
		fmt.Printf("key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
